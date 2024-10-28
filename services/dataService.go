package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type table struct {
	Id         string `json:"*id"`
	InsertedAt string `json:"*insertedAt"`
	Data       []any  `json:"data,omitempty"`
	Name       string `json:"name,omitempty"`
}

func Insert(elem string, database string, tableName string) {
	var elemData = []byte(elem)

	pathFile := fmt.Sprintf("exampleSheets/%s.json", database)

	// Read the sheet
	fileData, fileErr := os.ReadFile(pathFile)
	if fileErr != nil {
		fmt.Printf("Database '%s' not found.", database)
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

	var res []table

	// Extract the sheet
	json.Unmarshal([]byte(string(fileData)), &res)

	// Get the table
	filled := false
	for i := 0; i < len(res) && !filled; i++ {
		if res[i].Name == tableName {
			res[i].Data = append(res[i].Data, data)
			filled = true
		}
	}

	newData, err := json.MarshalIndent(res, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(pathFile, newData, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTable(name string, pathFile string) {
	// Read the sheet
	fileData, fileErr := os.ReadFile(pathFile)
	if fileErr != nil {
		log.Fatal(fileErr)
	}

	// Convert the data to structure
	data := make(map[string]interface{})

	// Set name
	data["name"] = name
	// Generate a default id
	data["*id"] = uuid.New()
	// Add date
	data["*insertedAt"] = time.Now().UTC().Format(time.RFC3339)
	data["data"] = []interface{}{}

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
