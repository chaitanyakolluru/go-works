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
		log.Fatal(err)
	}

	fileData := jsonFile.OpenFileAndReadData()
	if err := json.Unmarshal(fileData, &fileRecords); err != nil {
		log.Fatal(err)
	}

	fileRecords = append(fileRecords, record)

	fileRecordsJson, err := json.Marshal(fileRecords)
	if err != nil {
		log.Fatal(err)
	}

	jsonFile.WriteIntoFile(fileRecordsJson)
	c.IndentedJSON(http.StatusCreated, record)
}
