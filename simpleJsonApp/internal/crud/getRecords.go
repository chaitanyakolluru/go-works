package crud

import (
	"encoding/json"
	"log"
	"net/http"
	"simpleJsonApp/internal/jsonFile"

	"github.com/gin-gonic/gin"
)

// GetRecord godoc
//
//	@Summary		get all records
//
// @Schemes
//
//	@Description	gets all records within the json file
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	[]jsonFile.Record	"get all records"
//	@Failure		500		{string}	string	"internal server error"
//	@Router			/records [get]
func GetRecords(c *gin.Context) {
	var fileRecords []jsonFile.Record
	fileData := jsonFile.OpenFileAndReadData()
	if len(fileData) != 0 {
		if err := json.Unmarshal(fileData, &fileRecords); err != nil {
			log.Fatalf("cannot unmarshall file data, error: %s", err.Error())
		}

		c.IndentedJSON(http.StatusOK, fileRecords)
	} else {
		c.IndentedJSON(http.StatusOK, []jsonFile.Record{})
	}
}
