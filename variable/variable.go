package main

import (
	"fmt"
	"log"
)

func main() {
	var firstName string = "Singgih"
	var lastName = "Pratama"
	age := 25
	log.Println(firstName, lastName, age)

	fmt.Printf("Hallo %s %s %v! \n", firstName, lastName, age)

	/*
		Call func multiVariable
	*/
	multiVariable()
	predifinedVariable()
}

func multiVariable() {
	var firstName, middleName, lastName, age = "Muhammad", "Pratama", "Singgih", 20

	one, isFriday, ninePointSix, say := 1, true, 9.6, "wtf"

	log.Printf("Hello %s %s %s %v", firstName, middleName, lastName, age)
	fmt.Println(one, isFriday, ninePointSix, say)
}

func predifinedVariable() {
	var firstName, _ = "Farah", "kuyy"

	log.Printf("Hello %s", firstName)
}
