package releases

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
	controller := NewController()

	collectionGroup := rg.Group("/releases")
	{
		collectionGroup.GET("/", controller.getReleases)
	}

	memberGroup := rg.Group("/release/:id")
	{
		memberGroup.GET("/", controller.getReleaseByID)
	}
}
