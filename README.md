# ChessGoGame
Implementation of the well-known chess game written on golang language. University project contains GUI, AI and Server-client part.

### Implementation of the base game

Everybody knows what a base chess game is. If you are not sure, you can check [here](https://en.wikipedia.org/wiki/Chess).

### List of the extra functionalities

1. Possibility of playing agains AI or another online connected player.

2. Possibilities of revert your turn through the Undo button. Your oponent can Accept or Deny your revert. If you are playing against AI you will be able to revert anytime.

3. You will be able see your possible moves during the game.

### How to run the game

1. go get -v github.com/ivamilusheva/ChessGoGame/pkg/game

2. import the package in your main package

3. Initialize Board, Node and Game as configured below. The recursion level is being set in the Game constructor.

	var board game.Board
	
	areWeWithWhite := true
	
	board.New(areWeWithWhite)
	
	board.PrintMatrix()
	
	var currentNode game.Node
	
	currentNode.New(board, 0)
	
	var game game.Game
	
	game.New(currentNode, 4)
	
	game.PlayGameWithAI()

