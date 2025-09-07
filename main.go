package main

import (
	"fmt"
	"main/game"
)

func main() {
	seaBattle := game.NewSeaBattle()
	if !seaBattle.Start() {
		seaBattle.Stop()
	}

	fmt.Println(seaBattle.GetNamePlayers()) // проверка что игрок добавляется
}
