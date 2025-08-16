package main

import (
	"fmt"
	"os"
	startgame "seabattle2/startGame"
	"seabattle2/terminal"
)

func main() {
	App()

}

func App() {
	players := make(map[string]startgame.Player)
	var currentPlayer startgame.Player

	for {
		var input string
		fmt.Printf("1.Начать игру \n2.Завершить приложение\n")
		fmt.Scanln(&input)

		switch input {
		case "1":
			terminal.ClearTerminal()

			ok := startgame.Registarion(players, &currentPlayer)

			if !ok {
				continue
			}
			terminal.ClearTerminal()
			startgame.StartGame(&currentPlayer)

		case "2":
			terminal.ClearTerminal()

			os.Exit(0)
		default:
			terminal.ClearTerminal()
			continue
		}

	}

}
