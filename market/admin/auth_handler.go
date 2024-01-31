package admin

import (
	"platform/authorization/identity"
	"platform/http/actionresults"
	"platform/http/handling"
	"platform/sessions"
)

type AuthenticationHandler struct {
	identity.User
	identity.SignInManager
	identity.UserStore
	sessions.Session
	handling.URLGenerator
}

const SigningMsgKey string = "signing_message"

func (handler AuthenticationHandler) GetSignIn() actionresults.ActionResult {
	message := handler.Session.GetValueDefault(SigningMsgKey, "").(string)
	return actionresults.NewTemplateAction("signin.html", message)
}

type Credentials struct {
	Username string
	Password string
}

func (handler AuthenticationHandler) PostSignIn(creds Credentials) actionresults.ActionResult {
	if creds.Password == "mysecret" {
		user, ok := handler.UserStore.GetUserByName(creds.Username)

		if ok {
			handler.Session.SetValue(SigningMsgKey, "")
			handler.SignInManager.SignIn(user)
			return actionresults.NewRedirectAction("/admin/section/")
		}
	}

	handler.Session.SetValue(SigningMsgKey, "Access Denied")
	return actionresults.NewRedirectAction(mustGenerateUrl(handler.URLGenerator, AuthenticationHandler.GetSignIn))
}

func (handler AuthenticationHandler) PostSignOut(creds Credentials) actionresults.ActionResult {
	handler.SignInManager.SignOut(handler.User)

	return actionresults.NewRedirectAction("/")
}
