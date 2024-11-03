package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type table struct {
	Id         string                   `json:"*id"`
	InsertedAt string                   `json:"*insertedAt"`
	Data       []map[string]interface{} `json:"data,omitempty"`
	Name       string                   `json:"name,omitempty"`
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
	data["*id"] = uuid.New()
	// Add date
	data["*insertedAt"] = time.Now().UTC().Format(time.RFC3339)

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

func Select(selectors string, database string, tableName string, condition string) string {
	pathFile := fmt.Sprintf("exampleSheets/%s.json", database)

	// Read the sheet
	fileData, fileErr := os.ReadFile(pathFile)
	if fileErr != nil {
		fmt.Printf("Database '%s' not found.", database)
	}

	var res []table
	// Convert the data to structure
	json.Unmarshal([]byte(string(fileData)), &res)

	if tableName == "*" {
		if selectors == "*" {
			data, err := json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}
			return string(data)
		}
	}

	var tableTarget table

	// Get the table
	for i := 0; i < len(res) && len(tableTarget.Id) <= 0; i++ {
		if res[i].Name == tableName {
			tableTarget = res[i]
		}
	}

	if strings.Contains(selectors, "*") {
		data, err := json.Marshal(tableTarget.Data)
		if err != nil {
			log.Fatal(err)
		}
		return string(data)
	}

	var result []map[string]interface{}
	selects := strings.Split(selectors, ",")

	for i := 0; i < len(tableTarget.Data); i++ {
		item := make(map[string]interface{})
		for si := 0; si < len(selects); si++ {
			selectIndex := selects[si]
			if strings.Contains(strings.ToLower(selectIndex), "as") {
				parts := strings.Split(selectIndex, "as")
				indexItem := strings.ReplaceAll(parts[1], " ", "")
				propItem := strings.ReplaceAll(parts[0], " ", "")

				item[indexItem] = tableTarget.Data[i][propItem]

			} else {
				selectIndex = strings.ReplaceAll(selectIndex, " ", "")
				item[selectIndex] = tableTarget.Data[i][selectIndex]
			}

		}
		if len(item) > 0 {
			result = append(result, item)
		}
	}
	if len(result) > 0 {
		data, err := json.Marshal(result)
		if err != nil {
			log.Fatal(err)
			return "null"
		}
		return string(data)
	}

	return "null"

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
