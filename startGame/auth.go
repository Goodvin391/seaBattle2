package startgame

import (
	"fmt"
	"seabattle2/terminal"
	"time"
)

func ChekReg(users map[string]Player) (name string, res bool) {
	for {

		fmt.Println("Введите ваше имя: ")
		fmt.Scanln(&name)
		terminal.ClearTerminal()
		if name == "" {
			continue
		}
		break
	}
	_, ok := users[name]
	if !ok {
		res = false
		return
	}
	res = true
	return
}

// func setName(players map[string]Player, currentPlayer Player) {

// 	players[name] = currentPlayer

// }

func Registarion(users map[string]Player, currentPlayer *Player) bool {

	var input string

	name, ok := ChekReg(users)

	if ok {

		fmt.Printf("Добро пожаловать %s\n", name)

		time.Sleep(1 * time.Second)
		terminal.ClearTerminal()
		return true
	}

	if !ok {
		fmt.Printf("Вы не зарегистрированны - %s. Хотите зарегистрироваться?\n 1.Да\n 2.Нет\n ", name)
		for {
			fmt.Scanln(&input)
			terminal.ClearTerminal()
			switch input {
			case "1":

				currentPlayer.Name = name
				users[name] = Player{
					Name: name,
				}
				fmt.Println("Вы зарегистрированы!")

				time.Sleep(1 * time.Second)
				terminal.ClearTerminal()
				return true

			case "2":
				fmt.Println("Вы не зарегистрировались!")
				time.Sleep(1 * time.Second)
				terminal.ClearTerminal()

				return false
			default:
				terminal.ClearTerminal()
				fmt.Printf("Хотите зарегистрироваться?\n 1.Да\n 2.Нет\n ")
				fmt.Println("Введите число.")
			}

		}

	}
	return false

}
