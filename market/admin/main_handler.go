package admin

import (
	"platform/http/actionresults"
	"platform/http/handling"
)

var sectionNames = []string{"Products", "Categories", "Orders", "Database"}

type Handler struct {
	handling.URLGenerator
}

type TemplateContext struct {
	Sections       []string
	ActiveSection  string
	SectionUrlFunc func(string) string
}

func (handler Handler) GetSection(section string) actionresults.ActionResult {
	return actionresults.NewTemplateAction("admin.html", TemplateContext{
		Sections:      sectionNames,
		ActiveSection: section,
		SectionUrlFunc: func(sec string) string {
			sectionUrl, _ := handler.GenerateUrl(Handler.GetSection, sec)
			return sectionUrl
		},
	})
}
