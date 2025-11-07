package main

import (
	"ProfilePages/db"
	"ProfilePages/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	r := gin.Default()

	r.POST("/users", handler.CreateUser)
	r.GET("/users/:id", handler.GetUser)

	log.Println("running on 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
