package main

import (
	"simpleJsonApp/docs"
	"simpleJsonApp/internal/crud"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Simple Json App
//	@version		0.1.0
//	@description	This is a simple json app, apis of which will be fed into a crossplane provider.

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
// @accept 	json
// @produce	json

func main() {
	router := gin.Default()
	// docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "Simple Json App"
	router.POST("/record", crud.CreateRecord)
	router.GET("/records", crud.GetRecords)
	router.GET("/records/:name", crud.GetRecordsByName)
	router.PUT("/record", crud.UpdateRecord)
	router.DELETE("/record", crud.DeleteRecord)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run("localhost:8080")
}
