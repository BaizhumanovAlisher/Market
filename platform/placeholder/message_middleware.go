package placeholder

import (
	"errors"
	"io"
	"platform/config"
	"platform/pipeline"
	"platform/services"
	"platform/templates"
)

type SimpleMessageComponent struct {
	Message string
	config.Configuration
}

func (c *SimpleMessageComponent) ImplementsProcessRequestWithServices() {}

func (c *SimpleMessageComponent) Init() {
	c.Message = c.Configuration.GetStringDefault("main:message", "Default Message")
}

func (c *SimpleMessageComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	var cfg config.Configuration

	services.GetService(&cfg)

	msg, ok := cfg.GetString("main:message")

	if ok {
		io.WriteString(ctx.ResponseWriter, msg)
	} else {
		ctx.Error(errors.New("cannot find config setting"))
	}

	next(ctx)
}

func (c *SimpleMessageComponent) ProcessRequestWithServices(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext), executor templates.TemplateExecutor) {
	err := executor.ExecTemplate(ctx.ResponseWriter, "simple_message.html", c.Message)

	if err != nil {
		ctx.Error(err)
	} else {
		next(ctx)
	}
}
