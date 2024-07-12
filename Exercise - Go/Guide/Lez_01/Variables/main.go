package main

import "fmt"

func main() {
	//var variabile string
	variabile := "Hello World"
	fmt.Printf("Variabile: %q\n", variabile) // %q --> string

	marca := "Ferrari"
	fmt.Printf("Marca: %q\n", marca)

	var myAge, majorAge int
	myAge = 18
	majorAge = 18
	isAdult := myAge >= majorAge
	fmt.Printf("My Age is %v and the Major Age is %v\n", myAge, majorAge)
	fmt.Printf("Are you adult: %v\n", isAdult) // %v --> defautlt, senza trasformazioni

}
