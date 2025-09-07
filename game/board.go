package game

import (
	"fmt"
)

type Board interface {
	Render()
	PlaceShipsOnBoard()
}

type BoardGame struct {
	Ships []Ship
	Cells map[rune][]*CellOfBoard
}

func NewBoardGame() Board {
	res := BoardGame{
		Cells: make(map[rune][]*CellOfBoard),
	}

	for i := 'A'; i <= 'J'; i++ {
		rowOfCells := make([]*CellOfBoard, 0, 10)
		for j := 0; j < 10; j++ {
			newCell := NewCell()
			rowOfCells = append(rowOfCells, newCell)
		}
		res.Cells[i] = rowOfCells

	}

	return &res

}

func (b *BoardGame) PlaceShipsOnBoard() {
	for _, size := range sizeShips {
		newShip := NewShip()
		newShip.PlaceShipRandomly(size, b)
	}

}

func (b *BoardGame) Render() {
	// if b.Cells == nil {
	// 	fmt.Println("Кораболи не расположены")
	// 	time.Sleep(2 * time.Second)
	// 	terminal.ClearTerminal()
	// 	return
	// }

	fmt.Print("  ")
	for col := 0; col < 10; col++ {
		fmt.Printf("%c ", 'A'+col)

	}
	fmt.Println()

	for row := 1; row <= 10; row++ {
		fmt.Printf("%-2d", row)
		currentLetter := 'A'

		for col := 0; col < 10; col++ {

			currentRowLetter := b.Cells[currentLetter][row-1]
			if currentRowLetter.ship {
				fmt.Printf("%s ", "o")
			} else {
				fmt.Printf("%s ", ".")
			}

			currentLetter++
		}
		fmt.Println()
	}

	fmt.Println("Нажмите ENTER чтобы продолжить")
	fmt.Scanln()
	ClearTerminal()
}

func (b *BoardGame) PlaceShip(size int, direction string, cell *CellOfBoard) bool {
	currentCell := cell
	switch direction {
	case "letters":
		for i := 0; i < size; i++ {
			b.Cells[currentCell.coordinateLette][currentCell.coordinateNumber].ship = true
			currentCell.coordinateLette++
		}
		return true

	case "numbers":

		for i := 0; i < size; i++ {
			b.Cells[currentCell.coordinateLette][currentCell.coordinateNumber].ship = true
			currentCell.coordinateNumber++
		}
		return true
	default:
		return false
	}
}
