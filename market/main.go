package main

import (
	"market/admin"
	"market/admin/auth"
	"market/models/repo"
	"market/store"
	"market/store/cart"
	"platform/authorization"
	"platform/http"
	"platform/http/handling"
	"platform/pipeline"
	"platform/pipeline/basic"
	"platform/services"
	"platform/sessions"

	"sync"
)

func registerServices() {
	services.RegisterDefaultServices()
	repo.RegisterSqlRepositoryService()
	sessions.RegisterSessionService()
	cart.RegisterCartService()
	authorization.RegisterDefaultSignInService()
	authorization.RegisterDefaultUserService()
	auth.RegisterUserStoreService()
}

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		&sessions.SessionComponent{},

		authorization.NewAuthComponent(
			"admin",
			authorization.NewRoleCondition("Administrator"),
			admin.Handler{},
			admin.ProductsHandler{},
			admin.CategoriesHandler{},
			admin.OrdersHandler{},
			admin.DatabaseHandler{},
			admin.SignOutHandler{},
		).AddFallback("/admin/section/", "^/admin[/]?$"),

		handling.NewRouter(
			handling.HandlerEntry{Handler: store.ProductHandler{}},
			handling.HandlerEntry{Handler: store.CategoryHandler{}},
			handling.HandlerEntry{Handler: store.CartHandler{}},
			handling.HandlerEntry{Handler: store.OrderHandler{}},
			handling.HandlerEntry{Prefix: "api", Handler: store.RestHandler{}},
			handling.HandlerEntry{Handler: admin.AuthenticationHandler{}}).
			AddMethodAlias("/", store.ProductHandler.GetProducts, 0, 1).
			AddMethodAlias("/products[/]?[A-z0-9]*?", store.ProductHandler.GetProducts, 0, 1),
	)
}

func main() {
	registerServices()
	results, err := services.Call(http.Serve, createPipeline())

	if err == nil {
		(results[0].(*sync.WaitGroup)).Wait()
	} else {
		panic(err)
	}
}
