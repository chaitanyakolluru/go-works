package main

import (
	"github.com/chaitanyakolluru/go-works/simpleJsonApp/docs"
	"github.com/chaitanyakolluru/go-works/simpleJsonApp/internal/crud"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Simple Json App
// @version		0.1.0
// @description	This is a simple json app, apis of which will be fed into a crossplane provider.
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host		localhost:8081
// @accept 	json
// @produce	json
func main() {
	router := gin.Default()

	// Add CORS middleware, use cors.DefaultConfig() for default settings.
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	json := router.Group("/json")
	docs.SwaggerInfo.BasePath = "/json"
	docs.SwaggerInfo.Title = "Simple Json App"

	json.POST("/record", crud.CreateRecord)
	json.GET("/records", crud.GetRecords)
	json.GET("/records/:name", crud.GetRecordsByName)
	json.PUT("/record", crud.UpdateRecord)
	json.DELETE("/record", crud.DeleteRecord)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Use(cors.New(config))

	router.Run("localhost:8081")
}
