package releases

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (ctrl *Controller) getReleases(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from albums",
	})
}

func (ctrl *Controller) getReleaseByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello from release %s", c.Param("id")),
	})
}
