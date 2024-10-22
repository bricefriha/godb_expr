package main

import (
	"fmt"
	"godb/services"
)

func main() {
	fmt.Println("Engine started!!")

	// Create table
	var res = services.Execute(`
	CREATE TABLE animalList.animals`)

	// Insert
	// var res = services.Execute(`
	// INSERT INTO animalList.animals (label, legs)
	// VALUES ("Cows", 4)`)

	fmt.Println(res)
}
