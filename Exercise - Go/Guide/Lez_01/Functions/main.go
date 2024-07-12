package main

import (
	"errors"
	"fmt"
	"strings"
)

type player struct {
	name    string
	age     int
	country string
}

// function
func printPlayer(input player) {
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("Struct Player: %v\n", input)
	fmt.Printf("\t- Name: %q\n\t- Age: %d\n\t- Country: %q\n", input.name, input.age, input.country)
	fmt.Println(strings.Repeat("-", 60))
}

// function return values
var ErrDivideByZero = errors.New("cannot divide by zero")

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// method
//   - value receiver type 		--> func (p player) print() {}
//   - reference receiver type 	--> func (p *player) print() {}
func (p player) info() {
	fmt.Printf("Player name is %v and his age is %d. He's from %v\n", p.name, p.age, p.country)
	fmt.Println(strings.Repeat("-", 60))
}

func main() {
	player1 := player{
		name:    "Mario",
		age:     12,
		country: "Spain",
	}

	// function invocation
	printPlayer(player1)

	// method invocation
	player1.info()

	// function with return value
	num1, num2 := 5, 5
	result, err := divide(num1, num2)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println(strings.Repeat("-", 60))
		return
	}
	fmt.Printf("Result of division: %d\n", result)
	fmt.Println(strings.Repeat("-", 60))
}
