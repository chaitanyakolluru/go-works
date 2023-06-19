package main

import (
	"simpleJsonApp/internal/crud"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/record", crud.CreateRecord)
	router.GET("/records", crud.GetRecords)
	router.GET("/records/:name", crud.GetRecordsByName)
	router.PUT("/record", crud.UpdateRecord)

	router.Run("localhost:8080")
}
