package artists

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
	c := NewController()

	memberGroup := rg.Group("/artist/:id")
	{
		memberGroup.GET("/", c.GetAlbumByID)
	}
}
