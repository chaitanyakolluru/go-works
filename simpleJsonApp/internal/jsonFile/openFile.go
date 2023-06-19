package jsonFile

import (
	"fmt"
	"log"
	"os"
)

func getFileName() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("cannot get current working directory, error: %s", err.Error())
	}
	return fmt.Sprintf("%s/data.json", cwd)
}

func isFileExists() bool {
	_, err := os.Stat(getFileName())
	return err == nil
}

func OpenFileAndReadData() []byte {
	if !isFileExists() {
		_, err := os.Create(getFileName())
		if err != nil {
			log.Fatalf("cannot create data.json file, error: %s", err.Error())
		}
	}

	data, err := os.ReadFile(getFileName())
	if err != nil {
		log.Fatalf("cannot read file data, error: %s", err.Error())
	}
	return data
}
