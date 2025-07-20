package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/srntz/harp-api/cmd/api/controllers/v1/artists"
	"github.com/srntz/harp-api/cmd/api/controllers/v1/releases"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	g := rg.Group("/v1")
	{
		releases.RegisterRoutes(g)
		artists.RegisterRoutes(g)
	}
}
