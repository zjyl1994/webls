package handler

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	blackfriday "github.com/russross/blackfriday/v2"
	"github.com/spf13/viper"
	"github.com/zjyl1994/webls/config"
	"github.com/zjyl1994/webls/service"
	"github.com/zjyl1994/webls/util"
)

func listDirHandler(currentPath, realDiskPath string, c *gin.Context) {
	infos, hasREADME, err := service.ListDir(realDiskPath)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var readmeContent string
	if hasREADME {
		mdfile, err := ioutil.ReadFile(filepath.Join(realDiskPath, config.ReadmeFilename))
		if err == nil {
			readmeContent = string(blackfriday.Run(mdfile))
		}
	}
	var upperPath, urlPath string
	if currentPath != "/" {
		upperPath = filepath.Dir(currentPath)
		urlPath = currentPath
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"infos":    infos,
		"path":     currentPath,
		"urlpath":  urlPath,
		"uppath":   upperPath,
		"sitename": viper.GetString("sitename"),
		"markdown": readmeContent,
		"since":    util.CopyrightYear(viper.GetString("since")),
		"author":   viper.GetString("author"),
	})
}

func PortalHandler(c *gin.Context) {
	realpath := filepath.Join(viper.GetString("path"), c.Param("path"))
	if stat, err := os.Stat(realpath); err == nil {
		if stat.IsDir() {
			listDirHandler(c.Param("path"), realpath, c)
		} else {
			c.File(realpath)
		}
	} else if os.IsNotExist(err) {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}
}
