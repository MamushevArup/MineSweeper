package mid

import "fmt"

func Format(data string) {
	fmt.Println(data)
}
func ShowGameOverMessage() {
	Format("Game over")
}
func ShowWinMessage() {
	Format("Congratulations! You won!")
}
func UncoverCell(grid [][]Cell, row, column int) {
	helpUncover(grid, row, column)
}
func IsGameOver() bool {
	for _, row := range GameV.Grid {
		for _, cell := range row {
			if cell.IsBomb && !cell.IsCovered {
				return true
			}
		}
	}
	return false
}
func IsGameWon() bool {
	return GameV.Remaining == 0
}
