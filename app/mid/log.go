package mid

import (
	"fmt"
	"time"
)

func CountNeighboringBombs(grid [][]Cell, row, column int) int {
	return helpCountNeighbor(grid, row, column)
}

func PrintGame() {
	decor()
	marking()
	fmt.Printf("\nRemaining bombs: %d\n", GameV.Remaining)
	fmt.Printf("Elapsed time: %s\n", time.Since(StartTime))
}
