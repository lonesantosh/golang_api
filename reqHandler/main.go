package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("welcome in Rest api")

	server := gin.Default()

	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE
	server.Run(":8080")              //http://localhost:8080/events

}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"Welcome ": "Santosh"})
}
