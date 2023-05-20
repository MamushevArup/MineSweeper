package mid

import (
	"fmt"
	"strings"
	"time"
)

func CountNeighboringBombs(grid [][]Cell, row, column int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			r := row + i
			c := column + j
			if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c].IsBomb {
				count++
			}
		}
	}
	return count
}
func UncoverCell(grid [][]Cell, row, column int) {
	if row < 0 || row >= len(grid) || column < 0 || column >= len(grid[0]) {
		return
	}

	cell := &grid[row][column]

	if !cell.IsCovered || cell.IsBomb || cell.IsFlagged {
		return
	}

	cell.IsCovered = false
	game.Remaining--

	if cell.Value == 0 {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				UncoverCell(grid, row+i, column+j)
			}
		}
	}
}

func UncoverCell(grid [][]Cell, row, column int) {
	if row < 0 || row >= len(grid) || column < 0 || column >= len(grid[0]) {
		return
	}

	cell := &grid[row][column]

	if !cell.IsCovered || cell.IsBomb || cell.IsFlagged {
		return
	}

	cell.IsCovered = false
	game.Remaining--

	if cell.Value == 0 {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}
				UncoverCell(grid, row+i, column+j)
			}
		}
	}
}

// Function to check if the game is over
func IsGameOver() bool {
	for _, row := range game.Grid {
		for _, cell := range row {
			if cell.IsBomb && !cell.IsCovered {
				return true
			}
		}
	}
	return false
}

// Function to check if the game is won
func IsGameWon() bool {
	return game.Remaining == 0
}

// Function to print the current game state
func PrintGame() {
	// Clear the screen

	// Print the row numbers
	fmt.Print("   ")
	for i := 1; i <= game.Columns; i++ {
		fmt.Printf("%2d  ", i)
	}
	fmt.Println()

	// Print the top border
	fmt.Print("  +")
	for i := 1; i <= game.Columns; i++ {
		fmt.Print("---+")
	}
	fmt.Println()

	// Print the grid
	for i, row := range game.Grid {
		// Print the column number
		fmt.Printf("%2d|", i+1)

		// Print the cells
		for _, cell := range row {
			if cell.IsCovered {
				if cell.IsFlagged {
					fmt.Print(" F |")
				} else {
					fmt.Print(" X |")
				}
			} else {
				if cell.IsBomb {
					fmt.Print(" B |")
				} else {
					fmt.Printf(" %d |", cell.Value)
				}
			}
		}

		// Print the bottom border
		fmt.Println("\n  +", strings.Repeat("---+", game.Columns))
	}

	// Print the remaining bombs count and elapsed time
	fmt.Printf("\nRemaining bombs: %d\n", game.TotalBombs)
	fmt.Printf("Elapsed time: %s\n", time.Since(startTime))
}

// Function to show a game over message
func ShowGameOverMessage() {
	fmt.Println("Game Over!")
}

// Function to show a win message
func ShowWinMessage() {
	fmt.Println("Congratulations! You won!")
}
