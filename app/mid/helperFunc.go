package mid

import (
	"fmt"
	"strings"
)

func helpCountNeighbor(grid [][]Cell, row, column int) int {
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
func helpUncover(grid [][]Cell, row, column int) {
	if row < 0 || row >= len(grid) || column < 0 || column >= len(grid[0]) {
		return
	}

	cell := &grid[row][column]

	if !cell.IsCovered || cell.IsBomb || cell.IsFlagged {
		return
	}

	cell.IsCovered = false
	GameV.Remaining--

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
func decor() {
	Format("   ")
	for i := 1; i <= GameV.Columns; i++ {
		fmt.Printf("%2d  ", i)
	}
	fmt.Println()

	// Print the top border
	fmt.Print("  +")
	for i := 1; i <= GameV.Columns; i++ {
		fmt.Print("---+")
	}
	fmt.Println()
}
func marking() {
	for i, row := range GameV.Grid {
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
		fmt.Println("\n  +", strings.Repeat("---+", GameV.Columns))
	}
}
