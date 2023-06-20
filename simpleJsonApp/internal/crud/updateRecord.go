package crud

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/chaitanyakolluru/go-works/simpleJsonApp/internal/jsonFile"

	"github.com/gin-gonic/gin"
)

// UpdateRecord godoc
//
//	@Summary		update record
//
// @Schemes
//
//	@Description	update a record within the json file
//	@Accept			json
//	@Produce		json
//	@Param			record	body		jsonFile.Record		true	"record to be updated"
//	@Success		200		{object}	jsonFile.Record	"updated record"
//	@Failure		500		{string}	string	"internal server error"
//
//	@Router			/record [put]
func (client *Client) UpdateRecord(c *gin.Context) {
	var record jsonFile.Record
	var fileRecords []jsonFile.Record

	if err := c.BindJSON(&record); err != nil {
		log.Fatalf("cannot bind incoming request data, error: %s", err.Error())
	}

	fileData := jsonFile.OpenFileAndReadData()
	if len(fileData) != 0 {
		if err := json.Unmarshal(fileData, &fileRecords); err != nil {
			log.Fatalf("cannot unmarshall file data, error: %s", err.Error())
		}

		for index, rec := range fileRecords {
			if rec.Name == record.Name {
				fileRecords[index] = record
				fileRecords[index].Id = rec.Id

				fileRecordsJson, err := json.Marshal(fileRecords)
				if err != nil {
					log.Fatalf("cannot marshall data to json byte, error: %s", err.Error())
				}
				jsonFile.WriteIntoFile(fileRecordsJson)

				c.IndentedJSON(http.StatusOK, record)
				return
			}
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("requested record with name: %s not found", record.Name)})

}
