package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Prompt the user for the difficulty level
	level := start.PromptDifficultyLevel()

	// Create a new Minesweeper game
	game = NewGame(level)

	// Start the timer
	startTime = time.Now()

	// Print the initial game state
	PrintGame()

	// Start the game loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter the coordinates (ex: (4,5)): ")
		scanner.Scan()
		input := scanner.Text()
		ProcessInput(input)
		if IsGameWon() {
			ShowWinMessage()
			break
		}
		if IsGameOver() {
			break
		}
		PrintGame()
	}

	if err := scanner.Err(); err != nil {
		if _, err := fmt.Fprintln(os.Stderr, "error reading input:", err); err != nil {
			log.Println("Error working with Fprintln")
		}
	}
}
