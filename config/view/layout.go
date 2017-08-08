package view

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/qor/render"
)

type ViewContextKey string

var (
	View                   *render.Render
	ViewDefaultContextName ViewContextKey = "ViewContext"
	viewValueMap           map[string]string
)

func init() {
	View = render.New()

	View.FuncMapMaker = func(render *render.Render, req *http.Request, w http.ResponseWriter) template.FuncMap {
		funcMap := template.FuncMap{}

		funcMap["active_menu_class"] = func(s string) string {
			log.Println("Call active_menu_class")
			currentMenu := strings.Split(req.URL.Path, "/")[1]
			if s == currentMenu {
				return "active"
			} else {
				return ""
			}
		}

		return funcMap
	}

	htmlSanitizer := bluemonday.UGCPolicy()
	View.RegisterFuncMap("raw", func(str string) template.HTML {
		return template.HTML(htmlSanitizer.Sanitize(str))
	})

	viewValueMap = make(map[string]string)
	viewValueMap["test"] = "TestPassingValue"
}

func ViewValueMap() map[string]string {
	tempMap := make(map[string]string)
	for key, value := range viewValueMap {
		tempMap[key] = value
	}
	return tempMap
}
