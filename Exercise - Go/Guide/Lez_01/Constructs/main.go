package main

import "fmt"

func main() {
	// if construct
	x, y := 1, 2
	if y == 0 {
		fmt.Printf("Cannot divide by 0")
		return
	}
	division := x / y
	fmt.Printf("Division: %d / %d = %d\n\n", x, y, division) // %d --> numeri

	// for construct
	fmt.Print("Number from 0 to 9: \t")
	for i := 0; i <= 9; i++ {
		fmt.Printf("%d \t", i)
	}
}
