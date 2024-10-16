package services

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

func AddToSheet(name string, pathFile string) {
	fileData, fileErr := os.ReadFile(pathFile)

	if fileErr != nil {
		log.Fatal(fileErr)
	}
	var res []string
	// Read the example file
	json.Unmarshal([]byte(string(fileData)), &res)

	// Add an animal
	newList := append(res, name)

	newData, err := json.MarshalIndent(newList, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(pathFile, newData, os.ModePerm)
}
func Insert(elem string, pathFile string) {
	var elemData = []byte(elem)
	// Read the sheet
	fileData, fileErr := os.ReadFile(pathFile)
	if fileErr != nil {
		log.Fatal(fileErr)
	}

	// Convert the data to structure
	var data map[string]interface{}
	err := json.Unmarshal(elemData, &data)
	if err != nil {
		panic(err)
	}
	// Generate a default id
	data["€id"] = uuid.New()
	// Add date
	data["€insertedAt"] = time.Now().UTC().Format(time.RFC3339)

	var res []any

	// Extract the sheet
	json.Unmarshal([]byte(string(fileData)), &res)

	// Add the line
	newList := append(res, data)

	newData, err := json.MarshalIndent(newList, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(pathFile, newData, os.ModePerm)
}
