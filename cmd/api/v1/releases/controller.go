package releases

import "github.com/gin-gonic/gin"

func RegisterController(rg *gin.RouterGroup) {
	gMultiple := rg.Group("/releases")
	gSingle := rg.Group("/release")

	gMultiple.GET("/", Get)

	gSingle.GET("/:id", GetSingle)
}
