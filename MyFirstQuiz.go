package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")
	fmt.Println("Rule: 'This game has NEGATIVE marking also i.e. +2 and -1' ")

	fmt.Println("Enter your name:")

	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can play!")
	} else {
		fmt.Println("You cannot play")
		return
	}

	score := 0

	fmt.Println("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "RTX 3090" || answer+" "+answer2 == "rtx 3090" {
		fmt.Println("Correct!")
		score += 2
	} else {
		fmt.Println("Incorrect!")
		score--
	}

	fmt.Println("How many cores dose the Ryzen 9 3900x have? ")
	var cores int
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score += 2
	} else {
		fmt.Println("Incorrect!")
		score--
	}

	fmt.Printf("You scored %v out of 4. \n", score)
	percent := (score / 4) * 100
	fmt.Printf("You scored: %v%%.\n", percent)

}
