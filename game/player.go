package game

import "fmt"

type Player interface {
	AskName()
	GetName() string
	AddBoard(Board)
	GetBoard() Board
}

type HumanPlayer struct {
	Name  string
	Board Board
}

func NewHumanPlayer() Player {
	return &HumanPlayer{}
}

func (hp *HumanPlayer) AskName() {
	var name string

	for {

		fmt.Println("Введите ваше имя: ")
		fmt.Scanln(&name)

		ClearTerminal()
		if name == "" {
			continue
		}
		break
	}
	hp.Name = name
}

func (hp *HumanPlayer) GetName() string {
	return hp.Name
}

func (hp *HumanPlayer) AddBoard(board Board) {
	hp.Board = board
}
func (hp *HumanPlayer) GetBoard() Board {
	return hp.Board
}
