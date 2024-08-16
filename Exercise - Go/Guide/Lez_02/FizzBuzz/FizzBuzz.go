package main

import "strconv"

//Requirements:
//1. Write a FizzBuzz method that accepts a number as input and returns it as a String.
//2. For multiples of three return Fizz instead of the number
//3. For the multiples of five return Buzz
//4. For numbers that are multiples of both three and five return FizzBuzz

func FizzBuzz(input int) string {
	if input%3 == 0 && input%5 == 0 {
		return "Fizzbuzz"
	} else if input%3 == 0 {
		return "Fizz"
	} else if input%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(input) // convert number to string
	}
}

func main() {
	for i := 1; i <= 15; i++ {
		println(FizzBuzz(i))
	}

}
