package module

import "github.com/gin-gonic/gin"

func AttachModuleRouterGroup(r *gin.Engine) {
	moduleGroup := r.Group("/module")
	moduleGroup.Group("/example") // TODO: module loading logic
}
