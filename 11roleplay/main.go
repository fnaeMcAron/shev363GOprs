package main

import (
	"fmt"
	"rpg-server/game"
	"rpg-server/server"
)

func main() {

	fmt.Println("Выберите режим:")
	fmt.Println("1. PvE")
	fmt.Println("2. PvP")

	var choice int
	fmt.Scan(&choice)

	if choice == 1 {

		player := &game.Player{
			Name:     "Неро",
			HP:       100,
			Strength: 15,
		}

		enemy := &game.Enemy{
			Name:     "Гамбино",
			HP:       80,
			Strength: 10,
		}

		game.StartPvE(player, enemy)

	} else {

		server.Start()
	}
}
