package crud

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"simpleJsonApp/internal/jsonFile"

	"github.com/gin-gonic/gin"
)

func DeleteRecord(c *gin.Context) {
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
				fileRecords[index] = fileRecords[len(fileRecords)-1]
				fileRecords = fileRecords[:len(fileRecords)-1]

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
