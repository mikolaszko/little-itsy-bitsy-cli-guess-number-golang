package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	greetUser()
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(100)
	var playerGuess int
	fmt.Println("Try to guess!")
	fmt.Scan(&playerGuess)
	numberOfTries := 1
	for randomNumber != playerGuess {
		numberOfTries++
		if randomNumber < int(playerGuess) {
			fmt.Println("Oops, you are guessing too high! Try something smaller 🤏")
		} else if randomNumber > int(playerGuess) {
			fmt.Println("Gotta guess higher! 🚀")
		}
		fmt.Scan(&playerGuess)
	}
	fmt.Printf("You won! It only took you %v tries!\n", numberOfTries)
}

func greetUser() {
	fmt.Println("Hello! You are here to play a game 🎲")
	fmt.Println("Here are the rules: just have fun and try to guess a number between 1 and 100. Good luck!")
}
