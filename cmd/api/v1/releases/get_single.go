package releases

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSingle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello from release %s", c.Param("id")),
	})
}
