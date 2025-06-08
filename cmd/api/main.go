package main

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/srntz/harp-api/cmd/api/v1"
	"github.com/srntz/harp-api/internal/initializers"
)

func main() {
	initializers.LoadEnv()

	e := gin.Default()

	v1.RegisterRootController(e)

	e.Run()
}
