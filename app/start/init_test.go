package start

import (
	"testing"
)

type testCase struct {
	level              int
	expectedRows       int
	expectedColumns    int
	expectedTotalBombs int
}

func TestNewGame(t *testing.T) {
	testCases := []testCase{
		{level: 1, expectedRows: 9, expectedColumns: 9, expectedTotalBombs: 10},
		{level: 2, expectedRows: 16, expectedColumns: 16, expectedTotalBombs: 40},
		{level: 3, expectedRows: 30, expectedColumns: 30, expectedTotalBombs: 99},
		{level: 4, expectedRows: 0, expectedColumns: 0, expectedTotalBombs: 0},
		{level: -1, expectedRows: 0, expectedColumns: 0, expectedTotalBombs: 0},
	}

	for _, tc := range testCases {
		t.Run(getTestName(tc.level), func(t *testing.T) {
			game := NewGame(tc.level)

			if game.Rows != tc.expectedRows {
				t.Errorf("Expected %d rows, but got %d", tc.expectedRows, game.Rows)
			}
			if game.Columns != tc.expectedColumns {
				t.Errorf("Expected %d columns, but got %d", tc.expectedColumns, game.Columns)
			}
			if game.TotalBombs != tc.expectedTotalBombs {
				t.Errorf("Expected %d total bombs, but got %d", tc.expectedTotalBombs, game.TotalBombs)
			}
		})
	}
}

func getTestName(level int) string {
	switch level {
	case 1:
		return "TestNewGame_BeginnerLevel"
	case 2:
		return "TestNewGame_IntermediateLevel"
	case 3:
		return "TestNewGame_ExpertLevel"
	case 4:
		return "TestNewGame_CustomLevel"
	case -1:
		return "TestNewGame_NegativeLevel"
	default:
		return "UnknownTestLevel"
	}
}
