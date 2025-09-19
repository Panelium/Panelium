package ui

import (
	"panelium/core/internal/router/ui/ui_routes"

	"github.com/gin-gonic/gin"
)

func AttachUIRouterGroup(r *gin.Engine) {
	uiGroup := r.Group("/ui")
	uiGroup.GET("/", ui_routes.HandleIndexGET)
}
