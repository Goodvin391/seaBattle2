package game

import (
	"fmt"
)

type Game interface {
	Start() bool
	Stop()
	GetNamePlayers() []string
}

type SeaBattle struct {
	status      bool
	listPlayers map[string]Player
}

func NewSeaBattle() Game {
	return &SeaBattle{
		listPlayers: make(map[string]Player),
	}
}

func (s *SeaBattle) Start() bool {
	s.status = true

	fmt.Println("Запуск игры")

	for {

		var input string
		fmt.Printf("1.Начать игру \n2.Завершить приложение\n")
		fmt.Scanln(&input)
		switch input {
		case "1":
			player := NewHumanPlayer()
			player.AskName()
			s.AddPlayer(player)
			// mainMenu(player)
			// запуск меню
			newBoard := NewBoardGame()
			player.AddBoard(newBoard)
			s.listPlayers[player.GetName()].GetBoard().PlaceShipsOnBoard()

			s.listPlayers[player.GetName()].GetBoard().Render()

		case "2":
			return true // временно
			// s.Stop()
		}

	}

}
func (s *SeaBattle) AddPlayer(player Player) {
	//player.AskName()
	name := player.GetName()
	_, ok := s.listPlayers[name]
	if !ok {
		s.listPlayers[name] = player
		return
	}
	fmt.Println("Этот игрок уже добавлен")

}

func (s *SeaBattle) Stop() {
	fmt.Println("Завершение игры")
	s.status = false
	Exit()
}
func (s *SeaBattle) GetNamePlayers() []string {
	res := []string{}

	for name := range s.listPlayers {
		res = append(res, name)
	}

	return res
}
