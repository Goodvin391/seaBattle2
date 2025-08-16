package startgame

import (
	"math/rand"
	"time"
)

func newShip(player *Player, size int) {

	for {

		letter, number := randomCell()
		enouthCells := checkLocation(letter, number, size)
		if !enouthCells {
			continue
		}

		chooseDirect := randomDirect()

		switch chooseDirect {
		case "letters":
			freeCells := chekCells(player, letter, number, size, chooseDirect)
			if !freeCells {
				continue
			}
			placeShip(player, letter, number, size, chooseDirect)
			return
		case "numbers":
			freeCells := chekCells(player, letter, number, size, chooseDirect)
			if !freeCells {
				continue
			}

			placeShip(player, letter, number, size, chooseDirect)
			return
		}

	}

}

func chekCells(player *Player, letter rune, number int, size int, direction string) bool {
	switch direction {
	case "numbers":
		for i := 0; i < size; i++ {
			currentCell := player.Board.Cells[letter][number]
			if currentCell.ship {
				return false
			}
			number++
		}
		return true
	case "letters":
		for i := 0; i < size; i++ {
			choosenCell := player.Board.Cells[letter][number]
			if choosenCell.ship {
				return false
			}
			letter++
		}

		return true
	default:
		return false
	}

}

func placeShip(player *Player, letter rune, number int, size int, direction string) {
	switch direction {
	case "letters":
		for i := 0; i < size; i++ {
			currentCell := Cell{
				coordinateLette:  letter,
				coordinateNumber: number,
				ship:             true,
			}

			player.Board.Cells[letter][number] = currentCell
			letter++
		}
	case "numbers":
		for i := 0; i < size; i++ {
			currentCell := Cell{
				coordinateLette:  letter,
				coordinateNumber: number,
				ship:             true,
			}
			player.Board.Cells[letter][number] = currentCell
			number++
		}
	}

}

func checkLocation(letter rune, number int, size int) bool {

	if int(letter)+size > 74 {
		return false
	} else if number+size > 10 {
		return false
	}

	return true

}

func randomCell() (coordinateLette rune, coordinateNumber int) {

	min := 65 // 'A'
	max := 74 // 'J'
	rand.Seed(time.Now().UnixNano())
	coordinateNumber = rand.Intn(10)
	coordinateLette = rune(rand.Intn(max-min+1) + min)

	return
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

func newBoard() Board {
	res := Board{
		Cells: make(map[rune][]Cell),
	}
	for i := 'A'; i <= 'J'; i++ {
		var rowOfCells []Cell
		for j := 0; j < 10; j++ {
			rowOfCells = append(rowOfCells, Cell{})
		}
		res.Cells[i] = rowOfCells

	}
	return res

}
