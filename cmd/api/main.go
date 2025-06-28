package main

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/srntz/harp-api/cmd/api/controllers/v1"
	"github.com/srntz/harp-api/internal/initializers"
)

func main() {
	initializers.LoadEnv()

	e := gin.Default()

	rg := e.Group("")
	{
		v1.RegisterRoutes(rg)
	}

	e.Run()
}
