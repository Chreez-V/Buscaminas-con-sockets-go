package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Board struct {
	Width  int
	Height int
	Mines  int
	Grid   [][]Cell
}

func NewBoard(width, height, mines int) *Board {
	b := &Board{
		Width:  width,
		Height: height,
		Mines:  mines,
		Grid:   make([][]Cell, height),
	}
	for i := range b.Grid {
		b.Grid[i] = make([]Cell, width)
	}
	b.placeMines()
	b.calculateAdjacents()
	return b
}

func (b *Board) placeMines() {
	rand.Seed(time.Now().UnixNano())
	placed := 0
	for placed < b.Mines {
		x := rand.Intn(b.Width)
		y := rand.Intn(b.Height)
		if !b.Grid[y][x].HasMine {
			b.Grid[y][x].HasMine = true
			placed++
		}
	}
}

func (b *Board) calculateAdjacents() {
	dirs := []struct{ dx, dy int }{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0},           {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			count := 0
			for _, d := range dirs {
				ny, nx := y+d.dy, x+d.dx
				if nx >= 0 && nx < b.Width && ny >= 0 && ny < b.Height {
					if b.Grid[ny][nx].HasMine {
						count++
					}
				}
			}
			b.Grid[y][x].AdjacentMines = count
		}
	}
}

func (b *Board) Reveal(x, y int) bool {
	if x < 0 || x >= b.Width || y < 0 || y >= b.Height {
		return false
	}
	cell := &b.Grid[y][x]
	if cell.Revealed || cell.Flagged {
		return false
	}
	cell.Revealed = true
	return !cell.HasMine
}

func (b *Board) Print() {
	fmt.Println("  ", "0 1 2 3 4")
	for y, row := range b.Grid {
		fmt.Printf("%d ", y)
		for _, c := range row {
			if c.Flagged {
				fmt.Print("F ")
			} else if !c.Revealed {
				fmt.Print(". ")
			} else if c.HasMine {
				fmt.Print("* ")
			} else {
				fmt.Printf("%d ", c.AdjacentMines)
			}
		}
		fmt.Println()
	}
}

