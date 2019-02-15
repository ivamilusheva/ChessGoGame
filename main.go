package main

import (

//		"fmt"
	"ChessGoPorj/game"
		// "math/rand"
		// "time"
	//	"bufio"
	//	"os"

	//	"strings"
	//	"image/color"
)

//import "notnil/chess"
//import "notnil/chessimg"

//import "notnil/svgo"

func main() {
	var board game.Board
	board.New(true)
	board.PrintMatrix()
	var currentNode game.Node
	currentNode.New(board, 0)

	// Max recustion count: 4
	var game game.Game
	game.New(currentNode, 4)
	game.PlayerPlay()
	game.PlayGame()
}

/*



	// func new game / clean board
	var matrix[8][8] string;
	for i:=0; i<8;i++{
		for j:=0;j<8;j++{
			matrix[i][j] = "-";
		}
	}
	m := make(map[string]string)

	m["k"]= "♔"
	m["q"]= "♕"
	m["r"]= "♖"
	m["b"]= "♗"
	m["n"]= "♘"
	m["p"]= "♙"
	m["K"]= "♚"
	m["Q"]= "♛"
	m["R"]= "♜"
	m["B"]= "♝"
	m["N"]= "♞"
	m["P"]= "♟"

	// func print board
//	var abcd = make();
	fmt.Println("  A B C D E F G H");
	for i:=0; i<8;i++{
		fmt.Print(8 - i, " ");
		for j:=0;j<8;j++{
			fmt.Print(matrix[i][j], " ");
		}
		fmt.Print(i + 1, " ");
		fmt.Println();
	}
	fmt.Println("  H G F E D C B A");
	//fmt.Println(matrix);
	fmt.Println();

	for i:=0;i<len(startFEN);i++{
		x := m[string(startFEN[i])];
		if x != ""{		fmt.Print(x, " "); }
	}

	//var t rune=0x2659;
	/*var Figurines = []rune{
		'.', ',',
		0x2659, 0x265F,
		0x2658, 0x265E,
		0x2657, 0x265D,
		0x2656, 0x265C,
		0x2655, 0x265B,
		0x2654, 0x265A,
	}
	fmt.Println(Figurines)
	game := chess.NewGame()
        // generate moves until game is over
        for game.Outcome() == chess.NoOutcome {
            // select a random move
            moves := game.ValidMoves()
            move := moves[rand.Intn(len(moves))]
            game.Move(move)
        }
        // print outcome and game PGN
        fmt.Println(game.Position().Board().Draw())
        fmt.Printf("Game completed. %s by %s.\n", game.Outcome(), game.Method())
		fmt.Println(game.String())     */

/*	game := chess.NewGame()
moves := game.ValidMoves()

//input := bufio.NewReader(os.Stdin)
//	fmt.Print("Enter size of the board you want to play: ")
//	sizeOfBoard, _ := input.ReadString('\n')
//	sizeOfBoard = strings.TrimRight(sizeOfBoard, "\r\n")
fmt.Println(moves)
game.Move(moves[7])
moves = game.ValidMoves()
fmt.Println(moves)
game.Move(moves[7])
moves = game.ValidMoves()
fmt.Println(moves)     */
//fmt.Println(moves[7])
//var board chess.Board;
//	fmt.Println(board.Draw());
//	fmt.Println(game.Position().Board().Draw());
/*
// create file
f, err := os.Create("example.svg")
if err != nil {
	// handle error
}
defer f.Close()

// create board position
fenStr := "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq - 0 1"
pos := &chess.Position{}
if err := pos.UnmarshalText([]byte(fenStr)); err != nil {
	// handle error
}

// write board SVG to file
yellow := color.RGBA{255, 255, 0, 1}
mark := chessimg.MarkSquares(yellow, chess.D2, chess.D4)
if err := chessimg.SVG(f, pos.Board(), mark); err != nil {
	// handle error
}

//fmt.Println(game)
*/
