package main

import (
	"bufio"
	"fmt"
	"github.com/MamushevArup/minesweeper/app/mid"
	"github.com/MamushevArup/minesweeper/app/start"
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

	mid.GameV = start.NewGame(level)

	// Start the timer
	mid.StartTime = time.Now()

	// Print the initial game state
	mid.PrintGame()

	// Start the game loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter the coordinates (ex: (4,5)): ")
		scanner.Scan()
		input := scanner.Text()
		start.ProcessInput(input)
		if mid.IsGameWon() {
			mid.ShowWinMessage()
			break
		}
		if mid.IsGameOver() {
			break
		}
		mid.PrintGame()
	}

	if err := scanner.Err(); err != nil {
		if _, err := fmt.Fprintln(os.Stderr, "error reading input:", err); err != nil {
			log.Println("Error working with Fprintln")
		}
	}
}
