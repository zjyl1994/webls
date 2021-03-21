package main

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zjyl1994/webls/assets"
	"github.com/zjyl1994/webls/config"
	"github.com/zjyl1994/webls/handler"
)

var tplFunc = template.FuncMap{
	"htmlSafe": func(html string) template.HTML {
		return template.HTML(html)
	},
}

func main() {
	var err error
	if !config.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	if gin.IsDebugging() {
		router.SetFuncMap(tplFunc)
		router.LoadHTMLGlob("assets/templates/*.html")
	} else {
		router.SetHTMLTemplate(template.Must(template.New("").Funcs(tplFunc).ParseFS(assets.TemplateAssets, "templates/*.html")))
	}
	router.GET("/*path", handler.PortalHandler)
	fmt.Println("Webls start with", viper.GetString("path"), "at", viper.GetString("listen"))
	err = router.Run(viper.GetString("listen"))
	if err != nil {
		panic(err)
	}
}
