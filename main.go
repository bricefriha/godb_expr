package main

import (
	"fmt"
	"godb/services"
)

func main() {
	fmt.Println("Engine started!!")

	// Add an animal
	services.AddAnimal("cows", "exampleDocs/animalList.json")
}
