package main

import (
	"simpleJsonApp/internal/crud"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/record", crud.CreateRecord)

	router.Run("localhost:8080")
}
