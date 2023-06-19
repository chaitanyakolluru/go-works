package jsonFile

import (
	"fmt"
	"log"
	"os"
)

func getFileName() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/data.json", cwd)
}

func isFileExists() bool {
	isExists, err := os.Stat(getFileName())
	if err != nil {
		log.Fatal(err)
	}

	return isExists == nil
}

func OpenFileAndReadData() []byte {

	if !isFileExists() {
		_, err := os.Create(getFileName())
		if err != nil {
			log.Fatal(err)
		}
	}

	data, err := os.ReadFile(getFileName())
	if err != nil {
		log.Fatal(err)
	}
	return data
}
