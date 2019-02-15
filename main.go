package main

import (
	"github.com/ivamilusheva/ChessGoGame/pkg/game"
)

func main() {
	var board game.Board
	areWeWithWhite := true
	board.New(areWeWithWhite)
	board.PrintMatrix()
	var currentNode game.Node
	currentNode.New(board, 0)

	// Max recustion count: 4
	var game game.Game
	game.New(currentNode, 4)
	game.PlayGameWithAI()
}
