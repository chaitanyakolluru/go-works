package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "THIS IS WORKING!!")
}

func main() {
	router := gin.Default()
	router.GET("/status", showStatus)
	router.Run("localhost:12345")

	// go mod init makeapi
	// go mod tidy
	// more info:
	// https://golang.org/doc/tutorial/web-service-gin
	// curl http://localhost:12345/status
	// "THIS IS WORKING!!"
}
