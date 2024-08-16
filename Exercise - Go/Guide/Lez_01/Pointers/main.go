package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Repeat("#", 50))
	intVal := 42
	fmt.Printf("intVal     : %d\n", intVal)

	// memory address
	fmt.Printf("intVal     - memory address: %d\n", &intVal)
	fmt.Println(strings.Repeat("#", 50))

	// pointer
	// var intPointer *int
	intPointer := &intVal
	fmt.Printf("intPointer : %d\n", *intPointer) // deference operator
	fmt.Printf("intPointer - memory address: %v\n", intPointer)
	fmt.Println(strings.Repeat("#", 50))

	// changing value
	intVal = 3
	fmt.Printf("intVal     : %d\n", intVal)
	fmt.Printf("intPointer : %d\n", *intPointer)
	fmt.Println(strings.Repeat("#", 50))
}
