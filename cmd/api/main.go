package main

import (
	"log"

	"github.com/gin-gonic/gin"
	v1 "github.com/srntz/harp-api/cmd/api/controllers/v1"
	"github.com/srntz/harp-api/internal/infrastructure/db"
	"github.com/srntz/harp-api/internal/initializers"
)

func main() {
	initializers.LoadEnv()
	if _, err := db.GetPool(db.DefaultURL()); err != nil {
		log.Fatalf("Could not establish a connection with the database while initializing the server: %s", err.Error())
	}

	e := gin.Default()

	rg := e.Group("")
	{
		v1.RegisterRoutes(rg)
	}

	e.Run()
}
