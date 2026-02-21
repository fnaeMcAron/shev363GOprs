package game

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func StartPvE(player *Player, enemy *Enemy) {
	firstTurnPlayer := rand.Intn(2) == 0
	for player.IsAlive() && enemy.IsAlive() {

		if firstTurnPlayer {
			player.Hit(enemy)
			if enemy.IsAlive() {
				enemy.Hit(player)
			}
		} else {
			enemy.Hit(player)
			if player.IsAlive() {
				player.Hit(enemy)
			}
		}

		fmt.Println("Выберите куда бить (голова/тело/ноги):")
		fmt.Scan(&player.HitPart)

		fmt.Println("Выберите что блокировать (голова/тело/ноги):")
		var block string
		fmt.Scan(&block)
		player.Block(block)

		enemy.HitPart = RandomPart()
		enemy.Block(RandomPart())

		player.Hit(enemy)
		if enemy.IsAlive() {
			enemy.Hit(player)
		}

		fmt.Println("Здоровье игрока:", player.HP)
		fmt.Println("Здоровье противника:", enemy.HP)
	}

	if player.IsAlive() {
		fmt.Println("Ты победил")
	} else {
		fmt.Println("Враг победил")
	}
}
