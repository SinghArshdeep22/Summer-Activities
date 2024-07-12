package main

import "fmt"

// struct definition
type player struct {
	name    string
	age     int
	country string
}

func main() {
	// initialization
	player1 := player{
		name:    "Mario",
		age:     12,
		country: "Spain",
	}
	fmt.Printf("Struct Player: %v\n", player1)
	fmt.Printf("\tPlayer name: %q,\n\tPlayer age: %d,\n\tPlayer country: %q\n", player1.name, player1.age, player1.country)
}
