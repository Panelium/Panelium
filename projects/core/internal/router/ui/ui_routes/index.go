package ui_routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleIndexGET(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello world")
}
