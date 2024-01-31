package admin

import (
	"market/models"
	"platform/http/actionresults"
	"platform/http/handling"
)

type OrdersHandler struct {
	models.Repository
	handling.URLGenerator
}

func (handler OrdersHandler) GetData() actionresults.ActionResult {
	return actionresults.NewTemplateAction("admin_orders.html", struct {
		Orders      []models.Order
		CallbackUrl string
	}{
		Orders:      handler.Repository.GetOrders(),
		CallbackUrl: mustGenerateUrl(handler.URLGenerator, OrdersHandler.PostOrderToggle),
	})
}

func (handler OrdersHandler) PostOrderToggle(ref EditReference) actionresults.ActionResult {
	order := handler.Repository.GetOrder(ref.ID)
	order.Shipped = !order.Shipped

	handler.Repository.SetOrderShipped(&order)
	return actionresults.NewRedirectAction(mustGenerateUrl(handler.URLGenerator, Handler.GetSection, "Orders"))
}
