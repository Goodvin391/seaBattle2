package startgame

import (
	"fmt"
	"os"
	"seabattle2/terminal"
	"time"
)

var sizeShips = []int{4, 3, 3, 3, 2, 2, 2, 1, 1, 1, 1}

type Cell struct {
	coordinateLette  rune
	coordinateNumber int
	ship             bool
}

type Ship struct {
	Size int
}

type Board struct {
	Ships []Ship
	Cells map[rune][]Cell
}

type Player struct {
	Name  string
	Board Board
}

type Renderable interface {
	Render()
}

type ShipPlacer interface {
	PlaceShipsRandomly()
}

func (p *Player) PlaceShipsRandomly() {

	for _, value := range sizeShips {

		newShip(p, value)

	}

}
func (b *Board) Render() {
	if b.Cells == nil {
		fmt.Println("Кораболи не расположены")
		time.Sleep(2 * time.Second)
		terminal.ClearTerminal()
		return
	}

	fmt.Print("   ")
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
	terminal.ClearTerminal()
}

func StartGame(player *Player) {

	for {
		var input string
		fmt.Print("Выберете действие: \n1.Расположить корабли\n2.Показать поле с кораблями\n3.Выйти из игры\n4.Закрыть приложение\n")
		fmt.Scanln(&input)

		switch input {
		case "1":
			terminal.ClearTerminal()
			player.Board = newBoard()
			player.PlaceShipsRandomly()

			fmt.Println("Корабли расположены рандомно")
			time.Sleep(2 * time.Second)
			terminal.ClearTerminal()
		case "2":
			terminal.ClearTerminal()
			player.Board.Render()

		case "3":
			terminal.ClearTerminal()
			return
		case "4":
			terminal.ClearTerminal()
			os.Exit(0)
		default:
			terminal.ClearTerminal()
			continue
		}
	}

}
