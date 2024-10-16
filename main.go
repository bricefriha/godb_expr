package main

import (
	"fmt"
	"godb/services"
)

func main() {
	fmt.Println("Engine started!!")

	// Add an animal
	services.Insert([]byte(`
	{ 
		"id": "dwdeifhei23", 
		"label": "cows" 
	}`), "exampleDocs/animalList.json")
}
