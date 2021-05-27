// Created by: Ryan-Shaw-2
// Created on: May 2021
//
// This program multiples two numbers using while loops

package main

import "fmt"

func main() {
	// This function multiples two numbers using while loops
	var number1 int
	var number2 int
	var counter = 0
	var total = 0

	// input
	fmt.Println("This program multiples two numbers using while loops")
	fmt.Println()
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&number1)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&number2)
	fmt.Println()

	// process
	for counter < number1 {
		total += number2
		counter++
	}

	// output
	fmt.Println("The answer is", total)
}
