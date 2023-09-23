package main

import (
	"fmt"
	"log"
)

func main() {
	multiConst()
	multiConstOneLine()
}

func multiConst() {

	const (
		square        = "kotak"
		isToday  bool = false
		numeric  int8 = 1
		floatNum      = 1.1
	)

	log.Println(square, isToday, numeric, floatNum)
}

func multiConstOneLine() {
	const firstName, lastName = "Pratama", "Singgih"
	const age, high int16 = 25, 165

	const city, kodePost = "Depok", 16453

	fmt.Printf("Hello nama saya, %s %s "+"usia saya %v dan tinggi saya %v cm"+" dan saya tinggal di kota %s %v \n", firstName, lastName, age, high, city, kodePost)
}
