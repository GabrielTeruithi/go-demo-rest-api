package main

import (
	"github.com/gin-gonic/gin"
	"gteruithi.com/demo-rest-api/db"
	"gteruithi.com/demo-rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}
