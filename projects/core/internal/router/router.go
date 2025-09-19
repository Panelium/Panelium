package router

import (
	"panelium/core/internal/router/api"
	"panelium/core/internal/router/module"
	"panelium/core/internal/router/ui"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	api.AttachAPIRouterGroup(r)
	ui.AttachUIRouterGroup(r)
	module.AttachModuleRouterGroup(r)

	return r
}
