package main

import (
	"fmt"
)

func main() {
	var (
		name            string
		difficultyLevel int
	)
	level := []string{"Beginner", "Intermediate", "Expert"}
	fmt.Print("Enter your name: ")
	if _, err := fmt.Scan(&name); err != nil {
		return
	}
	for i := 0; i < len(level); i++ {
		fmt.Print(i+1, " -> "+level[i]+"\n")
	}
	fmt.Print("Enter the difficulty: ")
	if _, err := fmt.Scan(&difficultyLevel); err != nil {
		return
	}
	showField(createField(15, 15))
}

const ROWS, MINE = 8, 10

func createField(rows int, column int) [][]string {
	arr := make([][]string, rows)
	for i := 0; i < rows; i++ {
		arr[i] = make([]string, column)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < column; j++ {
			arr[i][j] = "+"
		}
	}
	return arr
}
func showField(arr [][]string) {
	fmt.Print("  ")
	for j := 1; j <= len(arr[0]); j++ {
		fmt.Printf("%2d ", j)
	}
	fmt.Println()
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%2d ", i+1)
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], "  ")
		}
		fmt.Println()
	}
}
func startGame() {

}
