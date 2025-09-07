package game

import (
	"math/rand"
	"time"
)

var sizeShips = []int{4, 3, 3, 3, 2, 2, 2, 1, 1, 1, 1}

type Ship interface {
	PlaceShipRandomly(int, *BoardGame)
}

type SeaShip struct {
	Size int
}

func NewShip() Ship {
	return &SeaShip{}
}

func (ss *SeaShip) PlaceShipRandomly(size int, board *BoardGame) {

	for {
		newCell := NewCell()
		newCell.InputRandomCoordinate()
		CheckEnouthCells := newCell.CheckLocation(size, board)

		if !CheckEnouthCells {
			continue
		}

		choosenDirect := randomDirect()

		switch choosenDirect {
		case "letters":
			possibilityOfLocation := newCell.CheckPossibilityOfLocation(size, choosenDirect, board)

			if !possibilityOfLocation {
				continue
			}

			if !board.PlaceShip(size, choosenDirect, newCell) {
				continue
			}

			return

		case "numbers":
			possibilityOfLocation := newCell.CheckPossibilityOfLocation(size, choosenDirect, board)

			if !possibilityOfLocation {
				continue
			}

			if !board.PlaceShip(size, choosenDirect, newCell) {
				continue
			}

			return
		}

	}

}

func randomDirect() string {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)

	if randomNumber >= 50 {
		return "letters"
	} else if randomNumber <= 49 {
		return "numbers"
	}
	return ""

}
