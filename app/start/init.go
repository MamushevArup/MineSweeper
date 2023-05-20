package start

import (
	"fmt"
	"github.com/MamushevArup/minesweeper/app/mid"
	"math/rand"
	"strconv"
	"strings"
)

func printIntro() {
	mid.Format("Choose difficulty level:")
	mid.Format("1. Beginner (9x9 grid, 10 bombs)")
	mid.Format("2. Intermediate (16x16 grid, 40 bombs)")
	mid.Format("3. Expert (30x30 grid, 99 bombs)")
}

func PromptDifficultyLevel() int {
	printIntro()
	var level int
	return levelIdent(level)
}
func levelIdent(level int) int {
	for {
		fmt.Print("Enter the difficulty level (1-3): ")
		_, err := fmt.Scanln(&level)
		if err == nil && level >= 1 && level <= 3 {
			break
		}
		mid.Format("Invalid input. Please enter a valid difficulty level.")
	}
	return level
}

func fillGrid() {

}
func NewGame(level int) *mid.Game {
	var rows, columns, totalBombs int

	switch level {
	case 1: // Beginner
		rows, columns, totalBombs = 9, 9, 10
	case 2: // Intermediate
		rows, columns, totalBombs = 16, 16, 40
	case 3: // Expert
		rows, columns, totalBombs = 30, 30, 99
	}

	game := &mid.Game{
		Rows:       rows,
		Columns:    columns,
		TotalBombs: totalBombs,
		Remaining:  rows*columns - totalBombs,
	}
	// Create the grid
	game.Grid = make([][]mid.Cell, rows)
	for i := 0; i < rows; i++ {
		game.Grid[i] = make([]mid.Cell, columns)
		for j := 0; j < columns; j++ {
			game.Grid[i][j] = mid.Cell{
				IsBomb:    false,
				IsCovered: true,
				IsFlagged: false,
				Value:     0,
			}
		}
	}

	// Place bombs randomly on the grid
	bombsPlaced := 0
	for bombsPlaced < totalBombs {
		row := rand.Intn(rows)
		column := rand.Intn(columns)
		if !game.Grid[row][column].IsBomb {
			game.Grid[row][column].IsBomb = true
			bombsPlaced++
		}
	}

	// Calculate the value of each cell (number of neighboring bombs)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if !game.Grid[i][j].IsBomb {
				game.Grid[i][j].Value = mid.CountNeighboringBombs(game.Grid, i, j)
			}
		}
	}

	return game
}
func ProcessInput(input string) {
	// Parse the row and column coordinates from the input
	coords := parseCoordinates(input)
	if coords == nil {
		mid.Format("Invalid input. Please enter the row and column coordinates in the format 'row,column'.")
		return
	}

	row, column := coords[0], coords[1]

	if row < 0 || row >= mid.GameV.Rows || column < 0 || column >= mid.GameV.Columns {
		mid.Format("Invalid input. The row and column coordinates are out of bounds.")
		return
	}

	cell := &mid.GameV.Grid[row][column]

	if cell.IsCovered {
		if cell.IsFlagged {
			cell.IsFlagged = false
		} else {
			cell.IsCovered = false
			if cell.IsBomb {
				mid.UncoverCell(mid.GameV.Grid, row, column)
				mid.ShowGameOverMessage()
			} else {
				mid.GameV.Remaining--
				if cell.Value == 0 {
					mid.UncoverCell(mid.GameV.Grid, row, column)
				}
			}
		}
	} else {
		mid.Format("Invalid input. The cell is already uncovered.")
	}
}
func parseCoordinates(key string) []int {
	if len(key) < 4 || key[0] != '(' || key[len(key)-1] != ')' {
		return nil
	}

	coords := make([]int, 2)
	coordsStr := key[1 : len(key)-1]
	coordsSplit := strings.Split(coordsStr, ",")
	if len(coordsSplit) != 2 {
		return nil
	}

	for i, coordStr := range coordsSplit {
		coord, err := strconv.Atoi(coordStr)
		if err != nil {
			return nil
		}
		coords[i] = coord - 1 // Adjust for zero-based indexing
	}

	return coords
}
