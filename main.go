package main

import (
	"encoding/json"
	"fmt"
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
	newAnimal := "cows"
	newList := append(res, newAnimal)

	newData, err := json.MarshalIndent(newList, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("exampleDocs/animalList.json", newData, os.ModePerm)

}
