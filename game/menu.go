package game

import "fmt"

func mainMenu(player Player) {
	for {
		var input string
		fmt.Print("Выберете действие: \n1.Расположить корабли\n2.Показать поле с кораблями\n3.Выйти из игры\n4.Закрыть приложение\n")
		fmt.Scanln(&input)

		switch input {
		case "1":
			ClearTerminal()
			// TO DO расположить корабли
			newBoard := NewBoardGame()
			player.AddBoard(newBoard)
			fmt.Println("Карта сгенерирована")

		case "2":
			player.GetBoard().Render()

		}

	}
}
