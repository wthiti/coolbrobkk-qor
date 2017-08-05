package config

import (
  "html/template"

  "github.com/microcosm-cc/bluemonday"
  "github.com/qor/render"
)

var (
  View *render.Render
)

func init() {
  View = render.New()

  htmlSanitizer := bluemonday.UGCPolicy()
  View.RegisterFuncMap("raw", func(str string) template.HTML {
    return template.HTML(htmlSanitizer.Sanitize(str))
  })
}
