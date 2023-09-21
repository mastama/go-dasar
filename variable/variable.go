package main

import (
	"fmt"
	"log"
)

func main() {
	var firstName string = "Singgih"
	var lastName = "Pratama"
	var age = 25
	log.Println(firstName, lastName, age)

	fmt.Printf("Hallo %s %s %v! \n", firstName, lastName, age)
}
