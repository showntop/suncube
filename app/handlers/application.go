package handlers

import (
	"html/template"
	"net/http"
	// "os"

	"github.com/qor/i18n/inline_edit"
	renderx "github.com/unrolled/render" // or "gopkg.in/unrolled/render.v1"
	// renderx "github.com/qor/render"

	"github.com/showntop/suncube/i18n"
	// "github.com/showntop/suncube/models"
)

var render *renderx.Render
var BindResult map[string]map[string]interface{} = map[string]map[string]interface{}{
	"Result": make(map[string]interface{}),
}

func AppendResult(key string, result interface{}) {
	BindResult["Result"][key] = result
}

func init() {
	render = renderx.New(renderx.Options{
		Directory: "app/views",
		Layout:    "layouts/application",
	})
}

func currentLocale(req *http.Request) string {
	locale := "en-US"
	if cookie, err := req.Cookie("locale"); err == nil {
		locale = cookie.Value
	}
	return locale
}

func I18nFuncMap(req *http.Request) template.FuncMap {
	return inline_edit.FuncMap(i18n.I18n, currentLocale(req), false)
}

// func render(req *http.Request, rw http.ResponseWriter, snippet template.HTML) {
// 	RootPath := os.Getenv("GOPATH") + "/src/github.com/showntop/suncube"

// 	tmpl2, _ := template.ParseFiles(RootPath + "/app/views/layouts/application.tmpl")

// 	tmpl2.Funcs(I18nFuncMap(req)).Execute(rw, map[string]map[string]interface{}{
// 		"Result": map[string]interface{}{
// 			"CurrentUser":     models.User{},
// 			"ContentTemplate": snippet,
// 			"CurrentLocale":   currentLocale(req),
// 		},
// 	})
// }
