package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Seed the random number generator for better randomness
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//rand.Seed(time.Now().UnixNano())

	// Generate a secret number between 1 and 100
	secretNumber := r.Intn(100) + 1

	// Welcome the user and explain the game
	fmt.Println("Welcome to the number guessing game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")

	for {
		// Get the user's guess
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your guess: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Parse the guess into an integer
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Sorry, this is a number guesser so I need an integer.")
			continue
		}

		// Check if the guess is correct
		if guess == secretNumber {
			fmt.Println("Congratulations!, you guessed the number!")
			break
		} else if guess < secretNumber {
			fmt.Println("Too low!")
		} else {
			fmt.Println("Too high!")
		}
	}
}
