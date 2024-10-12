package services

import (
	"encoding/json"
	"log"
	"os"
)

func AddAnimal(name string, pathFile string) {
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
