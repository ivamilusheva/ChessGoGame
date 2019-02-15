package game

type Node struct {
	board        Board
	currentValue int
}

func (n *Node) New(board Board, currentValue int) {
	n.board = board
	n.currentValue = currentValue
}
