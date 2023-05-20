package mid

import (
	"fmt"
	"time"
)

func CountNeighboringBombs(grid [][]Cell, row, column int) int {
	return helpCountNeighbor(grid, row, column)
}

// Function to print the current game state
func PrintGame() {
	// Print the row numbers
	decor()
	// Print the grid
	marking()
	// Print the remaining bombs count and elapsed time
	fmt.Printf("\nRemaining bombs: %d\n", GameV.Remaining)
	fmt.Printf("Elapsed time: %s\n", time.Since(StartTime))
}
