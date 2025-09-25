package ui_routes

import (
	"net/http"
	"panelium/core/internal/dashboard/views/layouts"
	"panelium/core/internal/dashboard/views/pages"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "", layouts.WithBase(pages.Index(), "Home", "homepage", true))
}

func Attach(r *gin.Engine) {
	uiGroup := r.Group("/")
	uiGroup.GET("/", Index)
}
