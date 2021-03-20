package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/zjyl1994/webls/assets"
	"github.com/zjyl1994/webls/config"
	"github.com/zjyl1994/webls/handler"
)

func main() {
	var err error
	if !config.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}
	config.LoadArgs()
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"htmlSafe": func(html string) template.HTML {
			return template.HTML(html)
		},
	})
	if gin.IsDebugging() {
		router.LoadHTMLGlob("assets/templates/*.html")
	} else {
		router.SetHTMLTemplate(template.Must(template.New("").ParseFS(assets.TemplateAssets, "templates/*.html")))
	}
	router.GET("/*path", handler.PortalHandler)
	err = router.Run(config.Listen)
	if err != nil {
		panic(err)
	}
}
