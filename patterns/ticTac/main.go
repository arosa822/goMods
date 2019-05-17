package main

import (
	"fmt"
)

type board struct {
	tiles [3][]string
}

type move struct {
	player string
	x      int
	y      int
}

//#### methods ####

// Generate an empty board
func (b *board) BuildEmptyBoard() {
	fmt.Println("clearing board...")
	spacer := []string{"-", "-", "-"}
	for i := 0; i < 3; i++ {
		for _, j := range spacer {
			b.tiles[i] = append(b.tiles[i], j)
		}
	}
}

// modify the board on turns
func (m *move) moveTo(b *board) {
	b.tiles[m.y][m.x] = m.player
}

// show the updated board
func (b board) showBoard() {
	fmt.Println("#### Game Board ####")
	for _, j := range b.tiles {
		fmt.Println(j)
	}
}

func (b board) takeATurn(player string, x, y int) {
	var turn = move{player: player, x: x, y: y}
	turn.moveTo(&b)
	b.showBoard()
}

// goroutine here place in channel possibly?
func (b board) scanForWinner() {
	// horizontal scan
	horiz := make([]string, 3)
	vert := make([]string, 3)
	for v := 0; v < 3; v++ {
		for i := 0; i < 3; i++ {
			vert[i] = b.tiles[v][i]
			if v < 1 {
				for j, k := range b.tiles[i] {
					horiz[j] = k
				}
				horiz_result := checkScan(horiz)
				fmt.Println(horiz_result)
			}
		}
		fmt.Println(checkScan(vert))

	}
}

func checkScan(a []string) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}

func main() {
	var gameBoard board
	gameBoard.BuildEmptyBoard()
	gameBoard.showBoard()

	//var turn = move{player: "x", x: 1, y: 2}

	gameBoard.takeATurn("x", 0, 0)
	gameBoard.takeATurn("o", 0, 1)
	gameBoard.takeATurn("o", 0, 2)

	gameBoard.takeATurn("o", 0, 0)
	gameBoard.takeATurn("o", 1, 0)
	gameBoard.takeATurn("o", 2, 0)

	gameBoard.scanForWinner()

}
