package admin

import (
	"market/models"
	"platform/http/actionresults"
	"platform/http/handling"
	"platform/sessions"
)

type ProductsHandler struct {
	models.Repository
	handling.URLGenerator
	sessions.Session
}

type ProductTemplateContext struct {
	Products []models.Product
	EditId   int
	EditUrl  string
	SaveUrl  string
}

const ProductEditKey string = "product_edit"

func (handler ProductsHandler) GetData() actionresults.ActionResult {
	return actionresults.NewTemplateAction("admin_products.html",
		ProductTemplateContext{
			Products: handler.GetProducts(),
			EditId:   handler.Session.GetValueDefault(ProductEditKey, 0).(int),
			EditUrl:  mustGenerateUrl(handler.URLGenerator, ProductsHandler.PostProductEdit),
			SaveUrl:  mustGenerateUrl(handler.URLGenerator, ProductsHandler.PostProductSave),
		})
}

type EditReference struct {
	ID int
}

func (handler ProductsHandler) PostProductEdit(ref EditReference) actionresults.ActionResult {
	handler.Session.SetValue(ProductEditKey, ref.ID)

	return actionresults.NewRedirectAction(mustGenerateUrl(handler.URLGenerator, Handler.GetSection, "Products"))
}

type ProductSaveReference struct {
	Id                int
	Name, Description string
	Category          int
	Price             float64
}

func (handler ProductsHandler) PostProductSave(p ProductSaveReference) actionresults.ActionResult {
	handler.Repository.SaveProduct(
		&models.Product{
			ID: p.Id, Name: p.Name, Description: p.Description,
			Category: &models.Category{ID: p.Category},
			Price:    p.Price,
		})

	handler.Session.SetValue(ProductEditKey, 0)
	return actionresults.NewRedirectAction(mustGenerateUrl(handler.URLGenerator, Handler.GetSection, "Products"))
}

func mustGenerateUrl(gen handling.URLGenerator, target interface{}, data ...interface{}) string {
	url, err := gen.GenerateUrl(target, data...)

	if err != nil {
		panic(err)
	}

	return url
}
