package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/srntz/harp-api/cmd/api/v1/releases"
)

// Registers all endpoints and router groups under '/v1'
func RegisterRootController(e *gin.Engine) {
	registerController(e.Group(""))
}

func registerController(rg *gin.RouterGroup) {
	g := rg.Group("/v1")

	releases.RegisterController(g)
}
