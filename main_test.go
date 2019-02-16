package main

import (
	"testing"
	"container/list"
	"reflect"
	"github.com/ivamilusheva/ChessGoGame/pkg/game"
)

func TestGetPieceValue(t *testing.T) {

	var board game.Board
	board.New(true)

	total := board.GetPieceValue("P")
	if total != 1 {
		t.Errorf("GetPieceValue was incorrect, got: %d, want: %d.", total, 1)
	}
	total = board.GetPieceValue("p")
	if total != -1 {
		t.Errorf("GetPieceValue was incorrect, got: %d, want: %d.", total, -1)
	}
	total = board.GetPieceValue("z")
	if total != 0 {
		t.Errorf("GetPieceValue was incorrect, got: %d, want: %d.", total, 0)
	}
}

func TestGetPrintValue(t *testing.T) {

	var board game.Board
	board.New(true)

	total := board.GetPrintValue("p")
	if total != "♟" {
		t.Errorf("GetPrintValue was incorrect, got: %v, want: %v.", total, "♟")
	}
	total = board.GetPrintValue("P")
	if total != "♙" {
		t.Errorf("GetPrintValue was incorrect, got: %v, want: %v.", total, "♙")
	}
	total = board.GetPrintValue("s")
	if total != "☐" {
		t.Errorf("GetPrintValue was incorrect, got: %v, want: %v.", total, "☐")
	}
}

func TestGetValue (t *testing.T){

	var board game.Board
	board.New(true)

	total := board.GetValue()
	if total != 0 {
		t.Errorf("GetValue was incorrect, got: %v, want: %v.", total, 0)
	}
}

func TestGetFillValue (t *testing.T) {
	var board game.Board
	board.New(true)

	total := board.GetFillValue("p", true)
	if total != "P" {
		t.Errorf("GetFillValue was incorrect, got: %v, want: %v.", total, "P")
	}
	total = board.GetFillValue("P", true)
	if total != "P" {
		t.Errorf("GetFillValue was incorrect, got: %v, want: %v.", total, "P")
	}
	total = board.GetFillValue("p", false)
	if total != "p" {
		t.Errorf("GetFillValue was incorrect, got: %v, want: %v.", total, "p")
	}

	total = board.GetFillValue("p", false)
	if total != "p" {
		t.Errorf("GetFillValue was incorrect, got: %v, want: %v.", total, "p")
	}

}

func TestIsFinished(t *testing.T){
	var board game.Board
	board.New(true)

	total := board.IsFinished()
	if total != false {
		t.Errorf("IsFinished was incorrect, got: %v, want: %v.", total, false)
	}

}

func TestClone(t *testing.T){
	var board game.Board
	board.New(true)

	total := board.Clone()
	if total != board {
		t.Errorf("Clone was incorrect, got: %v, want: %v.", total, board)
	}
}

func TestGetWhiteIndexes(t *testing.T){
	var board game.Board
	board.New(true)

	total := board.GetWhiteIndexes()
	list := list.New()
	var tup game.Tuple
	tup.New(6,0)
	list.PushBack(tup)
	tup.New(6,1)
	list.PushBack(tup)
	tup.New(6,2)
	list.PushBack(tup)
	tup.New(6,3)
	list.PushBack(tup)
	tup.New(6,4)
	list.PushBack(tup)
	tup.New(6,5)
	list.PushBack(tup)
	tup.New(6,6)
	list.PushBack(tup)
	tup.New(6,7)
	list.PushBack(tup)

	tup.New(7,0)
	list.PushBack(tup)
	tup.New(7,1)
	list.PushBack(tup)
	tup.New(7,2)
	list.PushBack(tup)
	tup.New(7,3)
	list.PushBack(tup)
	tup.New(7,4)
	list.PushBack(tup)
	tup.New(7,5)
	list.PushBack(tup)
	tup.New(7,6)
	list.PushBack(tup)
	tup.New(7,7)
	list.PushBack(tup)

	if !reflect.DeepEqual(total, list) {
		t.Errorf("GetWhiteIndexes was incorrect, got: %v, want: %v.", total, list)
	}
}
// GetBlackIndexes


func TestGetBlackIndexes(t *testing.T){
	var board game.Board
	board.New(true)

	total := board.GetBlackIndexes()
	list := list.New()
	var tup game.Tuple
	tup.New(0,0)
	list.PushBack(tup)
	tup.New(0,1)
	list.PushBack(tup)
	tup.New(0,2)
	list.PushBack(tup)
	tup.New(0,3)
	list.PushBack(tup)
	tup.New(0,4)
	list.PushBack(tup)
	tup.New(0,5)
	list.PushBack(tup)
	tup.New(0,6)
	list.PushBack(tup)
	tup.New(0,7)
	list.PushBack(tup)

	tup.New(1,0)
	list.PushBack(tup)
	tup.New(1,1)
	list.PushBack(tup)
	tup.New(1,2)
	list.PushBack(tup)
	tup.New(1,3)
	list.PushBack(tup)
	tup.New(1,4)
	list.PushBack(tup)
	tup.New(1,5)
	list.PushBack(tup)
	tup.New(1,6)
	list.PushBack(tup)
	tup.New(1,7)
	list.PushBack(tup)

	if !reflect.DeepEqual(total, list) {
		t.Errorf("GetBlackIndexes was incorrect, got: %v, want: %v.", total, list)
	}
}

func TestAreDifferentColor (t *testing.T) {

	var board game.Board
	board.New(true)

	total := board.AreDifferentColor(0 , 0, 7, 7)
	if total != true {
		t.Errorf("AreDifferentColor was incorrect, got: %v, want: %v.", total, true)
	}

	total = board.AreDifferentColor(0 , 0, 0, 7)
	if total != false {
		t.Errorf("AreDifferentColor was incorrect, got: %v, want: %v.", total, false)
	}

	total = board.AreDifferentColor(0 , 0, 4, 7)
	if total != false {
		t.Errorf("AreDifferentColor was incorrect, got: %v, want: %v.", total, false)
	}
	
}
