package jsonFile

import (
	"log"
	"os"
)

func WriteIntoFile(fileData []byte) {
	if err := os.WriteFile(getFileName(), fileData, 0644); err != nil {
		log.Fatal(err)
	}
}
