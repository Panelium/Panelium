package router

import (
	"net/http"
	"panelium/core/internal/dashboard/static"
	"panelium/core/internal/server/templ"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.StaticFS("/static", http.FS(static.Static))

	r.HTMLRender = &templ.Renderer{}

	return r
}
