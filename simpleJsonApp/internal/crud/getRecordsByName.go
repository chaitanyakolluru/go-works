package crud

import (
	"encoding/json"
	"log"
	"net/http"
	"simpleJsonApp/internal/jsonFile"

	"github.com/gin-gonic/gin"
)

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
