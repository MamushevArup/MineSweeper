package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Cell struct {
	IsBomb    bool
	IsCovered bool
	IsFlagged bool
	Value     int
}

type Game struct {
	Grid       [][]Cell
	Rows       int
	Columns    int
	TotalBombs int
	Remaining  int
}

var game *Game
var startTime time.Time

func main() {
	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Prompt the user for the difficulty level
	level := promptDifficultyLevel()

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

// Function to prompt the user for the difficulty level
func promptDifficultyLevel() int {
	fmt.Println("Choose difficulty level:")
	fmt.Println("1. Beginner (9x9 grid, 10 bombs)")
	fmt.Println("2. Intermediate (16x16 grid, 40 bombs)")
	fmt.Println("3. Expert (30x30 grid, 99 bombs)")

	var level int
	for {
		fmt.Print("Enter the difficulty level (1-3): ")
		_, err := fmt.Scanln(&level)
		if err == nil && level >= 1 && level <= 3 {
			break
		}
		fmt.Println("Invalid input. Please enter a valid difficulty level.")
	}

	return level
}

// Function to create a new Minesweeper game based on the difficulty level
func NewGame(level int) *Game {
	var rows, columns, totalBombs int

	switch level {
	case 1: // Beginner
		rows, columns, totalBombs = 9, 9, 10
	case 2: // Intermediate
		rows, columns, totalBombs = 16, 16, 40
	case 3: // Expert
		rows, columns, totalBombs = 30, 30, 99
	}

	game := &Game{
		Rows:       rows,
		Columns:    columns,
		TotalBombs: totalBombs,
		Remaining:  rows*columns - totalBombs,
	}

	// Create the grid
	game.Grid = make([][]Cell, rows)
	for i := 0; i < rows; i++ {
		game.Grid[i] = make([]Cell, columns)
		for j := 0; j < columns; j++ {
			game.Grid[i][j] = Cell{
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
				game.Grid[i][j].Value = CountNeighboringBombs(game.Grid, i, j)
			}
		}
	}

	return game
}

// Function to count the number of neighboring bombs for a given cell
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

// Function to process user input
func ProcessInput(input string) {
	// Parse the row and column coordinates from the input
	coords := ParseCoordinates(input)
	if coords == nil {
		fmt.Println("Invalid input. Please enter the row and column coordinates in the format 'row,column'.")
		return
	}

	row, column := coords[0], coords[1]

	if row < 0 || row >= game.Rows || column < 0 || column >= game.Columns {
		fmt.Println("Invalid input. The row and column coordinates are out of bounds.")
		return
	}

	cell := &game.Grid[row][column]

	if cell.IsCovered {
		if cell.IsFlagged {
			cell.IsFlagged = false
		} else {
			cell.IsCovered = false
			if cell.IsBomb {
				UncoverCell(game.Grid, row, column)
				ShowGameOverMessage()
			} else {
				game.Remaining--
				if cell.Value == 0 {
					UncoverCell(game.Grid, row, column)
				}
			}
		}
	} else {
		fmt.Println("Invalid input. The cell is already uncovered.")
	}
}

// Function to uncover a cell and its neighboring cells recursively
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

// Function to parse the row and column coordinates from a key string
func ParseCoordinates(key string) []int {
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

// Function to show a game over message
func ShowGameOverMessage() {
	fmt.Println("Game Over!")
}

// Function to show a win message
func ShowWinMessage() {
	fmt.Println("Congratulations! You won!")
}
