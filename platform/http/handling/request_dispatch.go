package handling

import (
	"fmt"
	"io"
	"net/http"
	"platform/http/actionresults"
	"platform/http/handling/params"
	"platform/pipeline"
	"platform/services"
	"reflect"
	"strings"
)

func NewRouter(handlers ...HandlerEntry) *RouterComponent {
	routes := generateRoutes(handlers...)

	var urlGen URLGenerator
	services.GetService(&urlGen)

	if urlGen == nil {
		services.AddSingleton(func() URLGenerator {
			return &routeUrlGenerator{routes: routes}
		})
	} else {
		urlGen.AddRoutes(routes)
	}

	return &RouterComponent{routes: routes}
}

type RouterComponent struct {
	routes []Route
}

func (rc *RouterComponent) Init() {}

func (rc *RouterComponent) ProcessRequest(context *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	for _, route := range rc.routes {
		if strings.EqualFold(context.Request.Method, route.httpMethod) {
			matches := route.expression.FindAllStringSubmatch(context.URL.Path, -1)

			if len(matches) > 0 {
				var rawParamVals []string

				if len(matches[0]) > 1 {
					rawParamVals = matches[0][1:]
				}

				err := rc.invokeHandler(route, rawParamVals, context)

				if err == nil {
					next(context)
				} else {
					context.Error(err)
				}

				return
			}
		}
	}

	context.ResponseWriter.WriteHeader(http.StatusNotFound)
}

func (rc *RouterComponent) invokeHandler(route Route, rawParams []string, context *pipeline.ComponentContext) error {
	paramVals, err := params.GetParametersFromRequest(context.Request, route.handlerMethod, rawParams)

	if err == nil {
		structVal := reflect.New(route.handlerMethod.Type.In(0))
		services.PopulateForContext(context.Context(), structVal.Interface())

		paramVals = append([]reflect.Value{structVal.Elem()}, paramVals...)
		result := route.handlerMethod.Func.Call(paramVals)

		if len(result) > 0 {
			if action, ok := result[0].Interface().(actionresults.ActionResult); ok {
				invoker := createInvokehandlerFunc(context.Context(), rc.routes)

				err = services.PopulateForContextWithExtras(context.Context(), action, map[reflect.Type]reflect.Value{
					reflect.TypeOf(invoker): reflect.ValueOf(invoker),
				})

				err = services.PopulateForContext(context.Context(), action)

				if err == nil {
					err = action.Execute(&actionresults.ActionContext{
						Context:        context.Context(),
						ResponseWriter: context.ResponseWriter,
					})
				}
			} else {
				io.WriteString(context.ResponseWriter, fmt.Sprint(result[0].Interface()))
			}
		}
	}

	return err
}
