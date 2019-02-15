package game

import (
	"bufio"
	"container/list"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Game struct {
	maxCount    int
	currentNode Node
}

type TupleBoard struct {
	item1 Board
	item2 Board
}

type TupleBoardValue struct {
	item1 Board
	item2 Board
	value int
}

func (g *Game) New(currentNode Node, maxCount int) {
	g.currentNode = currentNode
	g.maxCount = maxCount
}

func (g *Game) getRandomNumber(number int) int {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(number)[:number]
	return arr[0]
}

func (g *Game) Min(list *list.List) int {
	min_value := 10000
	for x := list.Front(); x != nil; x = x.Next() {
		current_value := x.Value.(TupleBoardValue).value
		if min_value > current_value {
			min_value = current_value
		}
	}

	return min_value
}

func (g *Game) GetAtIndex(list *list.List, index int) TupleBoardValue {
	count := 0
	for x := list.Front(); x != nil; x = x.Next() {
		current_value := x.Value.(TupleBoardValue)
		if count == index {
			return current_value
		}
		count++
	}

	// Not reachable
	return list.Front().Value.(TupleBoardValue)
}

func (g *Game) Max(list *list.List) int {
	max_value := -10000
	for x := list.Front(); x != nil; x = x.Next() {
		current_value := x.Value.(TupleBoardValue).value
		if max_value < current_value {
			max_value = current_value
		}
	}

	return max_value
}

func (g *Game) PlayGameWithAI() {
	var afterPlay TupleBoard

	for !g.currentNode.board.IsFinished() {

		if g.currentNode.board.areWeWithWhite && g.currentNode.board.isWhiteOnMove {
			g.PlayerPlay()
		} else if g.currentNode.board.areWeWithWhite && !g.currentNode.board.isWhiteOnMove {
			afterPlay = g.FindMin(g.currentNode, 0)
			fmt.Println("AI played")
			g.currentNode.board = afterPlay.item2
			g.currentNode.board.PrintMatrix()
		} else if !g.currentNode.board.areWeWithWhite && g.currentNode.board.isWhiteOnMove {
			afterPlay = g.FindMax(g.currentNode, 0)
			fmt.Println("AI played")
			g.currentNode.board = afterPlay.item2
			g.currentNode.board.PrintMatrix()
		} else {
			g.PlayerPlay()
		}
	}

	value := g.currentNode.board.GetValue()
	if value > 0 {
		fmt.Println("The winner is player")
	} else if value < 0 {
		fmt.Println("The winner is computer")
	} else {
		fmt.Println("The game ends with draw!")
	}
}

func (g *Game) PlayerPlay() {
	input := bufio.NewReader(os.Stdin)
	for true {
		fmt.Println("Please enter position for your fugure(row col) - separated with space")

		line, _ := input.ReadString('\n')
		line = strings.TrimRight(line, "\r\n")
		parts := strings.Split(line, " ")
		row, _ := strconv.Atoi(parts[0])
		col, _ := strconv.Atoi(parts[1])

		if !g.currentNode.board.IsInBoard(row, col) {
			fmt.Println("Move is outside the board. Try again.")
			continue
		}
		if g.currentNode.board.Matrix[row][col] == Empty {
			fmt.Println("You selected empty square. Try again.")
			continue
		}

		if g.currentNode.board.areWeWithWhite && IsLower(g.currentNode.board.Matrix[row][col]) {
			fmt.Println("You selected opponent figure. Try again.")
			continue
		} else if !g.currentNode.board.areWeWithWhite && IsUpper(g.currentNode.board.Matrix[row][col]) {
			fmt.Println("You selected opponent figure. Try again.")
			continue
		}

		moves := g.currentNode.board.GetCurrentPossibleMoves(row, col)
		if moves.Len() == 0 {
			fmt.Println("No available moves for this position. Try again.")
			continue
		}

		fmt.Println("Possible moves:")
		for possibleMove := moves.Front(); possibleMove != nil; possibleMove = possibleMove.Next() {
			fmt.Println(possibleMove.Value.(Move).endX, " ", possibleMove.Value.(Move).endY)
		}

		fmt.Println("Write coords (row col) separated with space")
		line, _ = input.ReadString('\n')
		line = strings.TrimRight(line, "\r\n")
		parts = strings.Split(line, " ")
		toX, _ := strconv.Atoi(parts[0])
		toY, _ := strconv.Atoi(parts[1])

		isValid := list.New()

		for x := moves.Front(); x != nil; x = x.Next() {
			if x.Value.(Move).endX == toX && x.Value.(Move).endY == toY {
				isValid.PushBack(x.Value)
			}
		}

		if isValid.Len() == 0 {
			fmt.Println("This move is not allowed. Try again.")
			continue
		}

		move := Move{row, col, toX, toY, false}
		g.currentNode.board.PerformMove(move)
		break
	}

	g.currentNode.board.PrintMatrix()
	fmt.Println("You just played. Please wait...")
}

func (g *Game) FindMin(node Node, counter int) TupleBoard {
	board := node.board
	if counter == g.maxCount {
		return TupleBoard{board, board}
	}

	childs := board.GetChildBoards()
	if childs.Len() == 0 {
		return TupleBoard{board, board}
	}

	if board.IsFinished() {
		// Game finished, no need to move more
		return TupleBoard{board, board}
	}

	// Recursion
	results := list.New()
	for child := childs.Front(); child != nil; child = child.Next() {
		newNode := Node{child.Value.(Board), 0}
		resultBoard := g.FindMax(newNode, counter+1)

		results.PushBack(TupleBoard{resultBoard.item1, child.Value.(Board)})
	}

	// Calculate values
	dataResults := list.New()
	for x := results.Front(); x != nil; x = x.Next() {
		t_result := x.Value.(TupleBoard)
		t_value := t_result.item1.GetValue()
		dataResults.PushBack(TupleBoardValue{t_result.item1, t_result.item2, t_value})
	}

	// Find best values
	var bestResultValue = g.Min(dataResults)
	bestBoards := list.New()
	for x := dataResults.Front(); x != nil; x = x.Next() {
		if x.Value.(TupleBoardValue).value == bestResultValue {
			bestBoards.PushBack(x.Value)
		}
	}

	bestBoardsCount := bestBoards.Len()
	if bestBoardsCount == 1 {
		// Only one best result
		x := bestBoards.Front()
		result := x.Value.(TupleBoardValue)
		return TupleBoard{result.item1, result.item2}
	} else {
		// Many best results, take randomly one
		var randomIndex = g.getRandomNumber(bestBoardsCount)
		var result = g.GetAtIndex(bestBoards, randomIndex)
		return TupleBoard{result.item1, result.item2}
	}
}

func (g *Game) FindMax(node Node, counter int) TupleBoard {
	board := node.board
	if counter == g.maxCount {
		return TupleBoard{board, board}
	}

	childs := board.GetChildBoards()
	if childs.Len() == 0 {
		return TupleBoard{board, board}
	}

	if board.IsFinished() {
		// Game finished, no need to move more
		return TupleBoard{board, board}
	}

	// Recursion
	results := list.New()
	for child := childs.Front(); child != nil; child = child.Next() {
		newNode := Node{child.Value.(Board), 0}
		resultBoard := g.FindMin(newNode, counter+1)

		results.PushBack(TupleBoard{resultBoard.item1, child.Value.(Board)})
	}

	// Calculate values
	dataResults := list.New()
	for x := results.Front(); x != nil; x = x.Next() {
		t_result := x.Value.(TupleBoard)
		t_value := t_result.item1.GetValue()
		dataResults.PushBack(TupleBoardValue{t_result.item1, t_result.item2, t_value})
	}

	// Find best values
	var bestResultValue = g.Max(dataResults)

	bestBoards := list.New()
	for x := dataResults.Front(); x != nil; x = x.Next() {
		if x.Value.(TupleBoardValue).value == bestResultValue {
			bestBoards.PushBack(x.Value)
		}
	}

	bestBoardsCount := bestBoards.Len()
	if bestBoardsCount == 1 {
		// Only one best result
		x := bestBoards.Front()
		result := x.Value.(TupleBoardValue)
		return TupleBoard{result.item1, result.item2}
	} else {
		// Many best results, take randomly one
		var randomIndex = g.getRandomNumber(bestBoardsCount)
		var result = g.GetAtIndex(bestBoards, randomIndex)
		return TupleBoard{result.item1, result.item2}
	}
}
