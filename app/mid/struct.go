package mid

import "time"

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

var GameV *Game
var StartTime time.Time
