package main

import (
	"github.com/gin-gonic/gin"
	"rest-api/db"
	"rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
