package handler

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/zjyl1994/webls/config"
	"github.com/zjyl1994/webls/service"
)

func listDirHandler(currentPath, realDiskPath string, c *gin.Context) {
	infos, err := service.ListDir(realDiskPath)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var upperPath string
	if currentPath != "/" {
		upperPath = filepath.Dir(currentPath)
	} else {
		currentPath = ""
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"infos":  infos,
		"path":   currentPath,
		"uppath": upperPath,
	})
}

func PortalHandler(c *gin.Context) {
	realpath := filepath.Join(config.DataDir, c.Param("path"))
	if stat, err := os.Stat(realpath); err == nil {
		if stat.IsDir() {
			listDirHandler(c.Param("path"), realpath, c)
		} else {
			c.File(realpath)
		}
	} else if os.IsNotExist(err) {
		_ = c.AbortWithError(http.StatusNotFound, err)
	} else {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}
}
