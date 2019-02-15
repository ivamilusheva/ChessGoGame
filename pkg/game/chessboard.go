package game

import (
	"strings"
	"math"
	"fmt"
	"container/list"
)

type Board struct {
	areWeWithWhite bool
	isWhiteOnMove  bool
	Matrix         [8][8]string
}

const (
	N     = 8
	Empty = "☐"
)

func (b *Board) New(areWeWithWhite bool) {
	b.isWhiteOnMove = true
	b.areWeWithWhite = areWeWithWhite
	b.FillEmptyMatrix()
}

func (b *Board) GetPieceValue(c string) int {
	var result int
	switch c {
	// пешка:
	case "P":
		result = 1
		break
	case "p":
		result = -1
		break
	// топ:
	case "R":
		result = 5
		break
	case "r":
		result = -5
		break
	// кон:
	case "N":
		result = 3
		break
	case "n":
		result = -3
		break
	// офицер:
	case "B":
		result = 3
		break
	case "b":
		result = -3
		break
	// царица:
	case "Q":
		result = 9
		break
	case "q":
		result = -9
		break
	// цар:
	case "K":
		result = 1000
		break
	case "k":
		result = -1000
		break
	default:
		result = 0
		break
	}

	if b.areWeWithWhite {
		return result
	} else {
		return result * -1
	}
}

func (b *Board) GetPrintValue(c string) string {
	switch c {
	// пешка:
case "P":
	return "♙"
case "p":
	return "♟"
// топ:
case "R":
	return "♖"
case "r":
	return "♜"
// кон:
case "N":
	return "♘"
case "n":
	return "♞"
// офицер:
case "B":
	return "♗"
case "b":
	return "♝"
// царица:
case "Q":
	return "♕"
case "q":
	return "♛"
// цар:
case "K":
	return "♔"
case "k":
	return "♚"
default:
		return Empty
	}
}
func (b *Board) GetValue() int {
	result := 0;
            for i := 0; i < N; i++ {
                for j := 0; j < N; j++ {
                    pieceValue := b.GetPieceValue(b.Matrix[i][j])
                    result += pieceValue;
                }
            }

            return result
}

func (b *Board) GetFillValue(c string, isWhite bool) string {

	if isWhite {
		return strings.ToUpper(c)
	} else {
		return c
	}
}

func (b *Board) FillEmptyMatrix() {
	for i := 2; i < N - 2; i++ {
		for j := 0; j < N; j++ {
			b.Matrix[i][j] = Empty;
		}
	}

	b.Matrix[0][0] = b.GetFillValue("r", !b.areWeWithWhite);
	b.Matrix[0][1] = b.GetFillValue("n", !b.areWeWithWhite);
	b.Matrix[0][2] = b.GetFillValue("b", !b.areWeWithWhite);
	b.Matrix[0][3] = b.GetFillValue("q", !b.areWeWithWhite);
	b.Matrix[0][4] = b.GetFillValue("k", !b.areWeWithWhite);
	b.Matrix[0][5] = b.GetFillValue("b", !b.areWeWithWhite);
	b.Matrix[0][6] = b.GetFillValue("n", !b.areWeWithWhite);
	b.Matrix[0][7] = b.GetFillValue("r", !b.areWeWithWhite);

	for j := 0; j < N; j++	{
		b.Matrix[1][j] = b.GetFillValue("p", !b.areWeWithWhite);
		b.Matrix[6][j] = b.GetFillValue("p", b.areWeWithWhite);
	}

	b.Matrix[7][0] = b.GetFillValue("r", b.areWeWithWhite);
	b.Matrix[7][1] = b.GetFillValue("n", b.areWeWithWhite);
	b.Matrix[7][2] = b.GetFillValue("b", b.areWeWithWhite);
	b.Matrix[7][3] = b.GetFillValue("q", b.areWeWithWhite);
	b.Matrix[7][4] = b.GetFillValue("k", b.areWeWithWhite);
	b.Matrix[7][5] = b.GetFillValue("b", b.areWeWithWhite);
	b.Matrix[7][6] = b.GetFillValue("n", b.areWeWithWhite);
	b.Matrix[7][7] = b.GetFillValue("r", b.areWeWithWhite);
}

func (b *Board) IsFinished() bool {
	value := b.GetValue()
    isFinished := int(math.Abs(float64(value))) > 500
    return isFinished
}

func (b *Board) Clone() Board {
	// b.areWeWithWhite=board.areWeWithWhite
	// b.isWhiteOnMove= board.isWhiteOnMove
	// b.Matrix= board.Matrix
	// return b

	var newMatrix [N][N]string

            for i := 0; i < N; i++ {
                for j := 0; j < N; j++ {
                    newMatrix[i][j] = b.Matrix[i][j];
                }
            }

            board := Board{b.areWeWithWhite, b.isWhiteOnMove, newMatrix }
return board
}

type Tuple struct {

	item1 int
	item2 int
}

func IsUpper(c string) bool {
	if strings.ToUpper(c) == c {
		return true
	} else {
		return false
	}

}

func (b *Board) GetWhiteIndexes() *list.List {
            //result := make ([]Tuple, 100)
			result := list.New()
			for i := 0; i < N; i++ {
                for j := 0; j < N; j++ {
                    pieceValue := b.Matrix[i][j]
                    if pieceValue != Empty && IsUpper(pieceValue) {
						var tuple Tuple;
						tuple.item1 = i
						tuple.item2 = j
						result.PushBack(tuple)
						//result = append(result, tuple)
                    }
                }
            }

			return result
}

func IsLower(c string) bool {
	if strings.ToLower(c) == c {
		return true
	} else {
		return false
	}

}

func (b *Board) GetBlackIndexes() *list.List {
	//var result = new List<Tuple<int, int>>();
	result := list.New() 
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			pieceValue := b.Matrix[i][j]
			if pieceValue != Empty && IsLower(pieceValue) {
				var tuple Tuple;
				tuple.item1 = i
				tuple.item2 = j
				result.PushBack(tuple)
			}
		}
	}

	return result
}

func (b *Board) AreDifferentColor(firstX, firstY, secondX, secondY int) bool {
            firstValue := b.Matrix[firstX][firstY]
            secondValue := b.Matrix[secondX][secondY]
            if IsUpper(firstValue) && IsLower(secondValue) {
                return true;
            }

            if IsLower(firstValue) && IsUpper(secondValue) {
                return true;
            }

            return false;
}

func (b *Board) IsInBoard(row, column int) bool {
	if row < 0 || row >= N	{
		return false;
	}

	if column < 0 || column >= N {
		return false;
	}

	return true;
}

func (b *Board) GetAllHorseFields(row, column int) *list.List {
	allMoves := list.New()
	allMoves.PushBack(Tuple { row + 1, column + 2});
	allMoves.PushBack(Tuple {row + 2, column + 1});
	allMoves.PushBack(Tuple {row - 1, column + 2});
	allMoves.PushBack(Tuple {row - 2, column + 1});
	allMoves.PushBack(Tuple {row + 1, column - 2});
	allMoves.PushBack(Tuple {row + 2, column - 1});
	allMoves.PushBack(Tuple {row - 1, column - 2});
	allMoves.PushBack(Tuple {row - 2, column - 1});
	result := list.New()
	for x := allMoves.Front(); x != nil; x = x.Next() {
		if b.IsInBoard(x.Value.(Tuple).item1, x.Value.(Tuple).item2) {
			result.PushBack(x.Value)
		}
	}
	//result := allMoves.Where(x => IsInBoard(x.Item1, x.Item2)).ToList();
	return result;
}

func (b *Board) GetAllKingFields(row, column int) *list.List {
	allMoves := list.New()
	allMoves.PushBack(Tuple {row + 1, column - 1});
		allMoves.PushBack(Tuple {row + 1, column});
		allMoves.PushBack(Tuple {row + 1, column + 1});
	
		allMoves.PushBack(Tuple {row, column - 1});
		allMoves.PushBack(Tuple {row, column + 1});
	
		allMoves.PushBack(Tuple {row - 1, column - 1});
		allMoves.PushBack(Tuple {row - 1, column});
		allMoves.PushBack(Tuple {row - 1, column + 1});
	
		result := list.New()
		for x := allMoves.Front(); x != nil; x = x.Next() {
			if b.IsInBoard(x.Value.(Tuple).item1, x.Value.(Tuple).item2) {
				result.PushBack(x.Value)
			}
		}
		return result;
}
        
func (b *Board) DoesOpponentLineBeatField(row, column int) bool {
	// надолу
	for i := row + 1; i < N; i++ {
		fieldValue := b.Matrix[i][column]
		if fieldValue == Empty	{
			continue;
		}

		// топ и царица
		blackCondition := (fieldValue == "q" || fieldValue == "r")
		whiteCondition := (fieldValue == "Q" || fieldValue == "R")
		if b.isWhiteOnMove && blackCondition {
			return true;
		} else if !b.isWhiteOnMove && whiteCondition	{
			return true;
		} else {
			// Друга фигура
			break;
		}
	}
	// нагоре
	for i := row - 1; i >= 0; i-- {
		fieldValue := b.Matrix[i][column]
		if fieldValue == Empty	{
			continue;
		}

		// топ и царица
		blackCondition := (fieldValue == "q" || fieldValue == "r")
		whiteCondition := (fieldValue == "Q" || fieldValue == "R")
		if b.isWhiteOnMove && blackCondition {
			return true;
		} else if !b.isWhiteOnMove && whiteCondition {
			return true;
		} else {
			// Друга фигура
			break;
		}
	}

	// Надясно
	for j := column + 1; j < N; j++	{
		fieldValue := b.Matrix[row][j]
		if fieldValue == Empty {
			continue;
		}

		// топ и царица
		blackCondition := (fieldValue == "q" || fieldValue == "r")
		whiteCondition := (fieldValue == "Q" || fieldValue == "R")
		if b.isWhiteOnMove && blackCondition {
			return true;
		} else if !b.isWhiteOnMove && whiteCondition {
			return true;
		} else {
			// Друга фигура
			break;
		}
	}

	// Наляво
	for j := column - 1; j >= 0; j-- {
		fieldValue := b.Matrix[row][j]
		if fieldValue == Empty {
			continue;
		}

		// топ и царица
		blackCondition := (fieldValue == "q" || fieldValue == "r")
		whiteCondition := (fieldValue == "Q" || fieldValue == "R")
		if b.isWhiteOnMove && blackCondition {
			return true;
		} else if !b.isWhiteOnMove && whiteCondition {
			return true;
		} else	{
			// Друга фигура
			break;
		}
	}

	return false;
}

func (b *Board) DoesOpponentDiagonalBeatField(row, column int) bool {
	// down right
	toX := row + 1
	toY := column + 1
	for toX < N && toY < N {
		if b.IsInBoard(toX, toY) {
			pieceValue := b.Matrix[toX][toY]
			if pieceValue != Empty	{
				// офицер и царица
				blackCondition := (pieceValue == "q" || pieceValue == "b")
				whiteCondition := (pieceValue == "Q" || pieceValue == "B")
				if b.isWhiteOnMove && blackCondition {
					return true;
				} else if !b.isWhiteOnMove && whiteCondition {
					return true;
				} else	{
					// Друга фигура
					break;
				}
			}
		} else	{
			// Извън дъската
			break;
		}

		toX++;
		toY++;
	}

	// down left
	toX = row + 1;
	toY = column - 1;
	for toX < N && toY >= 0	{
		if b.IsInBoard(toX, toY)	{
			pieceValue := b.Matrix[toX][toY]
			if pieceValue != Empty	{
				// офицер и царица
				blackCondition := (pieceValue == "q" || pieceValue == "b")
				whiteCondition := (pieceValue == "Q" || pieceValue == "B")
				if b.isWhiteOnMove && blackCondition {
					return true;
				} else if !b.isWhiteOnMove && whiteCondition {
					return true;
				}	else	{
					// Друга фигура
					break;
				}
			}
		}	else	{
			// Извън дъската
			break;
		}

		toX++;
		toY--;
	}

	// up right
	toX = row - 1;
	toY = column + 1;
	for toX >= 0 && toY < N	{
		if b.IsInBoard(toX, toY)	{
			pieceValue := b.Matrix[toX][toY]
			if pieceValue != Empty	{
				// офицер и царица
				blackCondition := (pieceValue == "q" || pieceValue == "b")
				whiteCondition := (pieceValue == "Q" || pieceValue == "B")
				if b.isWhiteOnMove && blackCondition {
					return true;
				}	else if !b.isWhiteOnMove && whiteCondition	{
					return true;
				}	else {
					// Друга фигура
					break;
				}
			}
		}	else	{
			// Извън дъската
			break;
		}

		toX--;
		toY++;
	}

	toX = row - 1;
	toY = column - 1;
	for toX >= 0 && toY >= 0 {
		if b.IsInBoard(toX, toY)	{
			pieceValue := b.Matrix[toX][toY]
			if pieceValue != Empty	{
				// офицер и царица
				blackCondition := (pieceValue == "q" || pieceValue == "b")
				whiteCondition := (pieceValue == "Q" || pieceValue == "B")
				if b.isWhiteOnMove && blackCondition {
					return true;
				} else if !b.isWhiteOnMove && whiteCondition {
					return true;
				}	else	{
					// Друга фигура
					break;
				}
			}
		}	else	{
			// Извън дъската
			break;
		}

		toX--;
		toY--;
	}

	return false;
}


func (b *Board) DoesOpponentHorseBeatField(row, column int) bool {
	
	allMoves := b.GetAllHorseFields(row, column)
	//result := list.New()
	for possibleMove := allMoves.Front(); possibleMove != nil; possibleMove = possibleMove.Next() {
		x := possibleMove.Value.(Tuple).item1
		y := possibleMove.Value.(Tuple).item2
		pieceValue := b.Matrix[x][y]
		if pieceValue != Empty	{
			// Кон
			blackCondition := (pieceValue == "n")
			whiteCondition := (pieceValue == "N")
			if b.isWhiteOnMove && blackCondition {
				return true;
			}	else if !b.isWhiteOnMove && whiteCondition	{
				return true;
			}	else	{
				// Друга фигура
				break;
			}
		}
	}

	return false;
}

func (b *Board) DoesOpponentPawnBeatField(row, column int) bool{
	if b.isWhiteOnMove {
		toX := row - 1

		// лява фигура дaли е черна пешка
		toY := column - 1
		if b.IsInBoard(toX, toY) && b.Matrix[toX][toY] == "p" {
			return true;
		}

		// дясна фигура дaли е черна пешка
		toY = column + 1
		if b.IsInBoard(toX, toY) && b.Matrix[toX][toY] == "p" {
			return true;
		}
	}	else	{
		toX := row + 1
		// лява фигура дaли е бяла пешка
		toY := column - 1
		if b.IsInBoard(toX, toY) && b.Matrix[toX][toY] == "P" {
			return true;
		}

		// дясна фигура дaли е бяла пешка
		toY = column + 1
		if b.IsInBoard(toX, toY) && b.Matrix[toX][toY] == "P" {
			return true;
		}
	}

	return false;
}

func (b *Board) DoesOpponentKingBeatField(row, column int) bool {
	allMoves := b.GetAllKingFields(row, column)
	for possibleMove := allMoves.Front(); possibleMove != nil; possibleMove = possibleMove.Next() {
		x := possibleMove.Value.(Tuple).item1
		y := possibleMove.Value.(Tuple).item2
		if b.IsInBoard(x, y) {
			pieceValue := b.Matrix[x][y]
			if pieceValue != Empty {
				// Цар
				blackCondition := (pieceValue == "k")
				whiteCondition := (pieceValue == "K")
				if b.isWhiteOnMove && blackCondition {
					return true;
				}	else if !b.isWhiteOnMove && whiteCondition {
					return true;
				}	else	{
					// Друга фигура
					break;
				}
			}
		}
	}

	return false;
}

func (b *Board) IsOpponentBeatField(row, column int) bool {
	var result = 
		b.DoesOpponentLineBeatField(row, column) ||
		b.DoesOpponentDiagonalBeatField(row, column) ||
		b.DoesOpponentHorseBeatField(row, column) ||
		b.DoesOpponentPawnBeatField(row, column) ||
		b.DoesOpponentKingBeatField(row, column);
	return result;
}

func (b *Board) GetDownPossibleMoves(row, column int) *list.List {
	result := list.New()
	for i := row + 1; i < N; i++ {
		if b.Matrix[i][column] == Empty	{
			result.PushBack(Move{row, column, i, column, false})
		}	else {
			if b.AreDifferentColor(row, column, i, column) {
				// взимане на фигура
				result.PushBack(Move{row, column, i, column, false})
			}
			break;
		}
	}
	return result;
}

func (b *Board) GetUpPossibleMoves(row, column int) *list.List {
	result := list.New()
	for i := row - 1; i >= 0; i-- {
		if b.Matrix[i][column] == Empty	{
			result.PushBack(Move{row, column, i, column, false})
		}	else	{
			if b.AreDifferentColor(row, column, i, column)	{
				// взимане на фигура
				result.PushBack(Move{row, column, i, column, false})
			}
			break;
		}
	}
	return result;
}

func (b *Board) GetLeftPossibleMoves(row, column int) *list.List {
	result := list.New()
	for j := column - 1; j >= 0; j-- {
		if b.Matrix[row][j] == Empty {
			result.PushBack(Move{row, column, row, j, false})
		}	else	{
			if b.AreDifferentColor(row, column, row, j)	{
				// взимане на фигура
				result.PushBack(Move{row, column, row, j, false})
			}
			break;
		}
	}
	return result;
}

func (b *Board) GetRightPossibleMoves(row, column int) *list.List {
	result := list.New()
	for j := column + 1; j < N; j++	{
		if b.Matrix[row][j] == Empty {
			result.PushBack(Move{row, column, row, j, false})
		}	else	{
			if b.AreDifferentColor(row, column, row, j)	{
				// взимане на фигура
				result.PushBack(Move{row, column, row, j, false})
			}
			break;
		}
	}
	return result;
}

func (b *Board) GetDiagonalDownRightPossibleMoves(row, column int) *list.List {
	result := list.New()

	toX := row + 1
	toY := column + 1
	for  toX < N && toY < N	{
		if b.IsInBoard(toX, toY)	{
			pieceValue := b.Matrix[toX][toY]
			if pieceValue == Empty	{
				result.PushBack(Move{row, column, toX, toY, false})
			}	else	{
				if b.AreDifferentColor(row, column, toX, toY) {
					// Взимане на фигура
					result.PushBack(Move{row, column, toX, toY, false})
				}
				break;
			}
		} else	{
			// Извън дъската
			break;
		}
		toX++
		toY++
	}
	return result
}

func (b *Board) GetDiagonalDownLeftPossibleMoves(row, column int) *list.List {
	result := list.New()

	toX := row + 1
	toY := column - 1
	for toX < N && toY >= 0	{
		if b.IsInBoard(toX, toY) {
			pieceValue := b.Matrix[toX][toY]
			if pieceValue == Empty	{
				result.PushBack(Move{row, column, toX, toY, false})
			} else	{
				if b.AreDifferentColor(row, column, toX, toY) {
					// Взимане на фигура
					result.PushBack(Move{row, column, toX, toY, false})
				}
				break;
			}
		} else	{
			// Извън дъската
			break;
		}
		toX++;
		toY--;
	}
	return result;
}

func (b *Board) GetDiagonalUpRightPossibleMoves(row, column int) *list.List {
	result := list.New()

	toX := row - 1
	toY := column + 1
	for toX >= 0 && toY < N	{
		if b.IsInBoard(toX, toY)	{
			pieceValue := b.Matrix[toX][toY]
			if pieceValue == Empty	{
				result.PushBack(Move{row, column, toX, toY, false})
			}	else {
				if b.AreDifferentColor(row, column, toX, toY) {
					// Взимане на фигура
					result.PushBack(Move{row, column, toX, toY, false})
				}
				break;
			}
		}	else {
			// Извън дъската
			break;
		}
		toX--;
		toY++;
	}
	return result;
}

func (b *Board) GetDiagonalUpLeftPossibleMoves(row, column int) *list.List {
	result := list.New()

	toX := row - 1
	toY := column - 1
	for toX >= 0 && toY >= 0 {
		if b.IsInBoard(toX, toY)	{
			pieceValue := b.Matrix[toX][toY]
			if pieceValue == Empty	{
				result.PushBack(Move{row, column, toX, toY, false})
			}	else	{
				if b.AreDifferentColor(row, column, toX, toY) {
					// Взимане на фигура
					result.PushBack(Move{row, column, toX, toY, false})
				}
				break;
			}
		} else {
			// Извън дъската
			break;
		}
		toX--;
		toY--;
	}
	return result;
}


func (b *Board) GetRockPossibleMoves(row int, column int) *list.List {
	result := list.New()
	
	downMoves := b.GetDownPossibleMoves(row, column)
        upMoves := b.GetUpPossibleMoves(row, column)
        leftMoves := b.GetLeftPossibleMoves(row, column)
        rightMoves := b.GetRightPossibleMoves(row, column)

	for temp := downMoves.Front(); temp != nil; temp = temp.Next() {
		result.PushBack(temp.Value)
	}
	for temp := upMoves.Front(); temp != nil; temp = temp.Next() {
		result.PushBack(temp.Value)
	}
	for temp := leftMoves.Front(); temp != nil; temp = temp.Next() {
		result.PushBack(temp.Value)
	}
	for temp := rightMoves.Front(); temp != nil; temp = temp.Next() {
		result.PushBack(temp.Value)
	}
	
	return result
}

func (b *Board) GetBishopPossibleMoves(row int, column int) *list.List {
	result := list.New()
	
	downRightMoves := b.GetDiagonalDownRightPossibleMoves(row, column)
        downLeftMoves := b.GetDiagonalDownLeftPossibleMoves(row, column)
        upRightMoves := b.GetDiagonalUpRightPossibleMoves(row, column)
        upLeftMoves := b.GetDiagonalUpLeftPossibleMoves(row, column)

	for temp := downRightMoves.Front(); temp != nil; temp = temp.Next() {
		result.PushBack(temp.Value)
	}
	for temp := downLeftMoves.Front(); temp != nil; temp = temp.Next() {
		result.PushBack(temp.Value)
	}
	for temp := upRightMoves.Front(); temp != nil; temp = temp.Next() {
		result.PushBack(temp.Value)
	}
	for temp := upLeftMoves.Front(); temp != nil; temp = temp.Next() {
		result.PushBack(temp.Value)
	}
	
	return result
}

func (b *Board) GetQueenPossibleMoves(row int, column int) *list.List {
	result := list.New()
	
	// line
        downMoves := b.GetDownPossibleMoves(row, column)
        upMoves := b.GetUpPossibleMoves(row, column)
        leftMoves := b.GetLeftPossibleMoves(row, column)
        rightMoves := b.GetRightPossibleMoves(row, column)

        // diagonal
        downRightMoves := b.GetDiagonalDownRightPossibleMoves(row, column)
        downLeftMoves := b.GetDiagonalDownLeftPossibleMoves(row, column)
        upRightMoves := b.GetDiagonalUpRightPossibleMoves(row, column)
        upLeftMoves := b.GetDiagonalUpLeftPossibleMoves(row, column)

	// line
	for temp := downMoves.Front(); temp != nil; temp = temp.Next() {
		result.PushBack(temp.Value)
	}
	for temp := upMoves.Front(); temp != nil; temp = temp.Next() {
		result.PushBack(temp.Value)
	}
	for temp := leftMoves.Front(); temp != nil; temp = temp.Next() {
		result.PushBack(temp.Value)
	}
	for temp := rightMoves.Front(); temp != nil; temp = temp.Next() {
		result.PushBack(temp.Value)
	}
	
	// diagonal
	for temp := downRightMoves.Front(); temp != nil; temp = temp.Next() {
		result.PushBack(temp.Value)
	}
	for temp := downLeftMoves.Front(); temp != nil; temp = temp.Next() {
		result.PushBack(temp.Value)
	}
	for temp := upRightMoves.Front(); temp != nil; temp = temp.Next() {
		result.PushBack(temp.Value)
	}
	for temp := upLeftMoves.Front(); temp != nil; temp = temp.Next() {
		result.PushBack(temp.Value)
	}
	
	return result
}

func (b *Board) GetHorsePossibleMoves(row int, column int) *list.List {
	allMoves := b.GetAllHorseFields(row, column)
	result := list.New()
	
	for possibleMove := allMoves.Front(); possibleMove != nil; possibleMove = possibleMove.Next() {
		x := possibleMove.Value.(Tuple).item1
		y := possibleMove.Value.(Tuple).item2
		pieceValue := b.Matrix[x][y]
		if pieceValue == Empty {
			var move Move
			move.New(row, column, x, y)
			result.PushBack(move)
		} else if b.AreDifferentColor(row, column, x, y) {
			var move Move
			move.New(row, column, x, y)
			result.PushBack(move)
		}
	}
	
	return result
}

func (b *Board) GetWhitePawnPossibleMoves(row int, column int) *list.List {
	result := list.New()

	// 1 позиция нагоре
	// Не може да излезне от дъската понеже се замества с друга фигура
	toX := row - 1
	toY := column
	isMovedToFinal := (toX == 0)
	if b.Matrix[toX][toY] == Empty {
		var move Move
		move.New(row, column, toX, toY)
		move.pawnReachedFinal = isMovedToFinal
		result.PushBack(move)
	}
	
	// Взимане на лява противникова фигура
	toY = column - 1
	if b.IsInBoard(toX, toY) && (b.Matrix[toX][toY] != Empty) && b.AreDifferentColor(row, column, toX, toY) {
		var move Move
		move.New(row, column, toX, toY)
		move.pawnReachedFinal = isMovedToFinal
		result.PushBack(move)
	}

	// Взимане на дясна противникова фигура
	toY = column + 1
	if b.IsInBoard(toX, toY) && (b.Matrix[toX][toY] != Empty) && b.AreDifferentColor(row, column, toX, toY) {
		var move Move
		move.New(row, column, toX, toY)
		move.pawnReachedFinal = isMovedToFinal
		result.PushBack(move)
	}

	// 2 позиции, ако не е играна досега
	if row == N - 2	{
		if b.Matrix[row - 1][column] == Empty && b.Matrix[row - 2][column] == Empty {
			var move Move
			move.New(row, column, row - 2, column)
			result.PushBack(move)
		}
	}
	
	return result
}

func (b *Board) GetBlackPawnPossibleMoves(row int, column int) *list.List {
	result := list.New()
	
	// 1 позиция надолу
	// Не може да излезне от дъската понеже се замества с друга фигура
	toX := row + 1
	toY := column
	isMovedToFinal := (toX == N - 1)
	if b.Matrix[toX][toY] == Empty {
		var move Move
		move.New(row, column, toX, toY)
		move.pawnReachedFinal = isMovedToFinal
		result.PushBack(move)
	}

	// Взимане на лява противникова фигура
	toY = column - 1
	if b.IsInBoard(toX, toY) && b.AreDifferentColor(row, column, toX, toY) {
		var move Move
		move.New(row, column, toX, toY)
		move.pawnReachedFinal = isMovedToFinal
		result.PushBack(move)
	}

	// Взимане на дясна противникова фигура
	toY = column + 1
	if b.IsInBoard(toX, toY) && b.AreDifferentColor(row, column, toX, toY) {
		var move Move
		move.New(row, column, toX, toY)
		move.pawnReachedFinal = isMovedToFinal
		result.PushBack(move)
	}

	// 2 позиции, ако не е играна досега
	if row == 1 {
		if b.Matrix[row + 1][column] == Empty && b.Matrix[row + 2][column] == Empty {
			var move Move
			move.New(row, column, row + 2, column)
			result.PushBack(move)
		}
	}

	return result
}

func (b *Board) GetPawnPossibleMoves(row int, column int) *list.List {
	pieceValue := b.Matrix[row][column]
	if IsUpper(pieceValue) {
		return b.GetWhitePawnPossibleMoves(row, column)
	} else {
		return b.GetBlackPawnPossibleMoves(row, column)
	}
}


func (b *Board) GetKingPossibleMoves(row int, column int) *list.List {
	allMoves := b.GetAllKingFields(row, column)
	result := list.New()
	
	for possibleMove := allMoves.Front(); possibleMove != nil; possibleMove = possibleMove.Next() {
		x := possibleMove.Value.(Tuple).item1
		y := possibleMove.Value.(Tuple).item2
		pieceValue := b.Matrix[x][y]

		isBeatenField := b.IsOpponentBeatField(row, column)
		if !isBeatenField {
			if pieceValue == Empty {
				var move Move
				move.New(row, column, x, y)
				result.PushBack(move)
			} else if b.AreDifferentColor(row, column, x, y) {
				// Взимане на фигура, няма break понеже ходовете на царя са точно определени
				var move Move
				move.New(row, column, x, y)
				result.PushBack(move)
			}
		}
	}
	
	return result
}

func (b *Board) GetCurrentPossibleMoves(x int, y int) *list.List {
	pieceValue := b.Matrix[x][y]
	if strings.ToLower(pieceValue) == "r" {
		// топ:
		return b.GetRockPossibleMoves(x, y)
	}
	
	if strings.ToLower(pieceValue) == "n" {
		// кон:
		return b.GetHorsePossibleMoves(x, y)
	}
	
	if strings.ToLower(pieceValue) == "b" {
		// офицер:
		return b.GetBishopPossibleMoves(x, y)
	}
	
	if strings.ToLower(pieceValue) == "q" {
		// царица:
		return b.GetQueenPossibleMoves(x, y)
	}
	
	if strings.ToLower(pieceValue) == "k" {
		// цар:
		return b.GetKingPossibleMoves(x, y)
	}
	
	if strings.ToLower(pieceValue) == "p" {
		// пешка:
		return b.GetPawnPossibleMoves(x, y)
	}
	
	return list.New()
}

func (b *Board) GetAllPossibleMoves() *list.List {
	indexes := list.New()
	if b.isWhiteOnMove {
		indexes = b.GetWhiteIndexes()
	} else {
		indexes = b.GetBlackIndexes()
	}
	
	result := list.New()
	for currentPosition:= indexes.Front(); currentPosition != nil; currentPosition = currentPosition.Next() {
		x := currentPosition.Value.(Tuple).item1
		y := currentPosition.Value.(Tuple).item2
		moves := b.GetCurrentPossibleMoves(x, y)
		for temp := moves.Front(); temp != nil; temp = temp.Next() {
			result.PushBack(temp.Value)
		}
	}
	
	return result
}

func (b *Board) GetChildBoards() *list.List {
	result := list.New()
	isGameFinished := b.IsFinished()
	if (isGameFinished) {
		return result
	}
	possibleMoves := b.GetAllPossibleMoves()
	for possibleMove := possibleMoves.Front(); possibleMove != nil; possibleMove = possibleMove.Next() {
		newBoard := b.Clone()
		newBoard.PerformMove(possibleMove.Value.(Move))
		newBoard.isWhiteOnMove = !b.isWhiteOnMove
		result.PushBack(newBoard)
	}

	return result
}

func (b *Board) PerformMove(move Move) {
	currentValue := b.Matrix[move.startX][move.startY]
	b.Matrix[move.startX][move.startY] = Empty
	if move.pawnReachedFinal {
		var queen string
		if b.isWhiteOnMove {
			queen = "Q"
		} else {
			queen = "q"
		}
		
		b.Matrix[move.endX][move.endY] = queen
	} else {
		b.Matrix[move.endX][move.endY] = currentValue
	}

	b.isWhiteOnMove = !b.isWhiteOnMove     
}

func (b *Board) PrintMatrix() {
	 for i := 0; i<N; i++ {
		 for j := 0; j<N; j++ {
			 
			 fmt.Print(b.GetPrintValue(b.Matrix[i][j]), " ")
		 }
		 fmt.Println();
	 }

}