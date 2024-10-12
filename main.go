package main

import (
	"encoding/json"
	"fmt"
	"godb/services"
	"log"
	"os"
)

func main() {
	fmt.Println("Engine started!!")
	fileData, fileErr := os.ReadFile("exampleDocs/animalList.json")

	if fileErr != nil {
		log.Fatal(fileErr)
	}
	var res []string
	// Read the example file
	json.Unmarshal([]byte(string(fileData)), &res)

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

	// Add an animal
	services.AddAnimal("cows", "exampleDocs/animalList.json")
}
