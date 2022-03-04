package main

import (
	"cfttest/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := router.Engine()
	r.Use(gin.Logger())
	if err := r.Run(":9081"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}
