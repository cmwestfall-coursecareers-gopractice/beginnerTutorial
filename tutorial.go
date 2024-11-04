package main

import "fmt"

func main () {
	fmt.Println("Welcome to my quiz game!")

	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	var age uint
	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can play!")
	} else {
		fmt.Println("Sorry, you can't play")
		return
	}

	score := 0

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer + " " + answer2 == "RTX 3090" {
		fmt.Println("Correct")
		score++;
	} else {
		fmt.Println("Incorrect")
	}
	
	fmt.Printf("You scored %v out of 1.", score)
}