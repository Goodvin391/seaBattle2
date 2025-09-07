package game

import (
	"math/rand"
	"time"
)

type Cell interface {
	GetRandomCell()
	InputRandomCoordinate()
}

type CellOfBoard struct {
	coordinateLette  rune
	coordinateNumber int
	ship             bool
}

func NewCell() *CellOfBoard {

	return &CellOfBoard{}
}

func (cb *CellOfBoard) InputRandomCoordinate() {
	min := 65 // 'A'
	max := 74 // 'J'
	rand.Seed(time.Now().UnixNano())
	cb.coordinateNumber = rand.Intn(10)
	cb.coordinateLette = rune(rand.Intn(max-min+1) + min)

}

func (cb *CellOfBoard) CheckLocation(size int, board *BoardGame) bool {

	letter := cb.coordinateLette
	number := cb.coordinateNumber
	if int(letter)+size > 74 {
		return false
	} else if number+size > 10 {
		return false
	}

	return true

}

func (cb *CellOfBoard) CheckPossibilityOfLocation(size int, direct string, board *BoardGame) bool {
	currentCell := *cb

	switch direct {
	case "letters":
		for i := 0; i < size; i++ {
			shipExists := board.Cells[currentCell.coordinateLette][currentCell.coordinateNumber].ship

			if shipExists {
				return false
			}
			// проверка соседних клеток
			if !ChekingNeughboringCells(&currentCell, board) {
				return false
			}

			currentCell.coordinateLette++

		}

		return true

	case "numbers":
		for i := 0; i < size; i++ {
			shipExists := board.Cells[currentCell.coordinateLette][currentCell.coordinateNumber].ship
			if shipExists {
				return false
			}
			if !ChekingNeughboringCells(&currentCell, board) {
				return false
			}
			currentCell.coordinateNumber++
		}

		return true

	default:
		return false

	}

}

func ChekingNeughboringCells(currentCell *CellOfBoard, board *BoardGame) bool {

	letter := currentCell.coordinateLette
	number := currentCell.coordinateNumber

	for i := -1; i < 2; i++ {

		for b := -1; b < 2; b++ {
			if letter+rune(i) < 65 || letter+rune(i) > 74 || number+b < 0 || number+b > 9 {

				continue
			}

			if board.Cells[letter+rune(i)][number+b].ship {
				return false
			}
		}

	}

	return true

}
