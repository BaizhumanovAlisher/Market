package cart

import (
	"encoding/json"
	"market/models"
	"platform/services"
	"platform/sessions"
	"strings"
)

const Key string = "cart"

func RegisterCartService() {
	services.AddScoped(func(session sessions.Session) Cart {
		var lines []*Line

		sessionVal := session.GetValue(Key)
		if strVal, ok := sessionVal.(string); ok {
			json.NewDecoder(strings.NewReader(strVal)).Decode(&lines)
		}

		return &sessionCart{
			BasicCart: &BasicCart{lines: lines},
			Session:   session,
		}
	})
}

type sessionCart struct {
	*BasicCart
	sessions.Session
}

func (sc *sessionCart) AddProduct(p models.Product) {
	sc.BasicCart.AddProduct(p)
	sc.SaveToSession()
}

func (sc *sessionCart) RemoveLineForProduct(id int) {
	sc.BasicCart.RemoveLineForProduct(id)
	sc.SaveToSession()
}

func (sc *sessionCart) SaveToSession() {
	builder := strings.Builder{}
	json.NewEncoder(&builder).Encode(sc.lines)
	sc.Session.SetValue(Key, builder.String())
}

func (sc *sessionCart) Reset() {
	sc.lines = []*Line{}
	sc.SaveToSession()
}
