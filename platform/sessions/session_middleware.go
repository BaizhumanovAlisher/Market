package sessions

import (
	"context"
	gorilla "github.com/gorilla/sessions"
	"platform/config"
	"platform/pipeline"
	"time"
)

type SessionComponent struct {
	store *gorilla.CookieStore
	config.Configuration
}

func (sc *SessionComponent) Init() {
	cookie, found := sc.Configuration.GetString("sessions:key")

	if !found {
		panic("Session key not found in configuration")
	}

	if sc.GetBoolDefault("sessions:cyclekey", true) {
		cookie += time.Now().String()
	}

	sc.store = gorilla.NewCookieStore([]byte(cookie))
}

func (sc *SessionComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	session, _ := sc.store.Get(ctx.Request, SessionContextKey)

	c := context.WithValue(ctx.Request.Context(), SessionContextKey, session)
	ctx.Request = ctx.Request.WithContext(c)

	next(ctx)
	session.Save(ctx.Request, ctx.ResponseWriter)
}
