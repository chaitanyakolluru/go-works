package crud

import (
	"encoding/json"
	"log"
	"net/http"
	"simpleJsonApp/internal/jsonFile"

	"github.com/gin-gonic/gin"
)

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

	fileRecords = append(fileRecords, record)

	fileRecordsJson, err := json.Marshal(fileRecords)
	if err != nil {
		log.Fatalf("cannot marshall data to json byte, error: %s", err.Error())
	}

	jsonFile.WriteIntoFile(fileRecordsJson)
	c.IndentedJSON(http.StatusCreated, record)
}
