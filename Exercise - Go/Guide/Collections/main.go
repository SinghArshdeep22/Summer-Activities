package main

import "fmt"

func main() {
	// 3 Collections --> Array | Slice | Map

	// arrays - non dynamic
	var grades [3]int
	grades[0] = 8
	grades[1] = 9
	grades[2] = 10

	fmt.Printf("Array Grades : %v\n", grades)

	// slice - dynamic
	friends := []string{"rick", "maggie"}
	friends = append(friends, "daryl")

	fmt.Printf("Slice Friends: %v\n", friends)
	for _, value := range friends {
		fmt.Printf("\tFriend %q\n", value)
	}

	// map - key/value - dictionary
	students := make(map[string]string)
	students["Cristiano"] = "5IA"
	students["Arsh"] = "5IA"

	fmt.Printf("Map Students: \n")
	for key, value := range students {
		fmt.Printf("\tStudente %q belongs to the %q\n", key, value)
	}
}
