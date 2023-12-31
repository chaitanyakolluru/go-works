package crud

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/chaitanyakolluru/go-works/simpleJsonApp/internal/jsonFile"

	"github.com/gin-gonic/gin"
)

// CreateRecord godoc
//
//	@Summary		create record
//
// @Schemes
//
//	@Description	create a record within the json file
//	@Accept			json
//	@Produce		json
//	@Param			record	body		jsonFile.Record		true	"record to be created"
//	@Success		201		{object}	jsonFile.Record	"created record"
//	@Failure		400		{string}	string	"failure message saying resource already exists"
//	@Failure		500		{string}	string	"internal server error"
//	@Router			/record [post]
//
// @Security ApiKeyAuth
func CreateRecord(c *gin.Context) {
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
	}

	if !jsonFile.ValidateRecord(fileRecords, record) {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("requested name: %s already exists", record.Name)})
		return
	}

	fileRecords = append(fileRecords, record)

	fileRecordsJson, err := json.Marshal(fileRecords)
	if err != nil {
		log.Fatalf("cannot marshall data to json byte, error: %s", err.Error())
	}

	jsonFile.WriteIntoFile(fileRecordsJson)
	c.IndentedJSON(http.StatusCreated, record)
}
