package crud

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/chaitanyakolluru/go-works/simpleJsonApp/internal/jsonFile"

	"github.com/gin-gonic/gin"
)

// GetRecordByName godoc
//
//	@Summary		get a single record
//
// @Schemes
//
//	@Description	gets a record within the json file
//	@Accept			json
//	@Produce		json
//	@Param			name	path			string		true	"name"
//	@Success		200		{object}		jsonFile.Record	"get a records"
//	@Failure		500		{string}	string	"internal server error"
//	@Router			/records/{name}	[get]
//
// @Security ApiKeyAuth
func GetRecordsByName(c *gin.Context) {
	name := c.Param("name")

	var fileRecords []jsonFile.Record
	fileData := jsonFile.OpenFileAndReadData()
	if len(fileData) != 0 {
		if err := json.Unmarshal(fileData, &fileRecords); err != nil {
			log.Fatalf("cannot unmarshall file data, error: %s", err.Error())
		}

		for _, record := range fileRecords {
			if record.Name == name {
				c.IndentedJSON(http.StatusOK, record)
				return
			}
		}
	}
	c.IndentedJSON(http.StatusOK, []jsonFile.Record{})
}
