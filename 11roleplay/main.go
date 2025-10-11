package main

import (
	"fmt"
	"math/rand"
)

type Character interface {
	Hit(target Character, bodyPart string)
	Block(bodyPart string)
}

type Player struct {
	Name        string
	HP          float32
	Strength    int
	blockedPart string
}

type Enemy struct {
	Name        string
	HP          float32
	Strength    int
	blockedPart string
}

func (p *Player) Hit(target Character, bodyPart string) {
	damage := float32(p.Strength)
	switch bodyPart {
	case "голова":
		damage *= 2
	case "тело":
		damage *= 1
	case "ноги":
		damage *= 1.5
	}
	switch t := target.(type) {
	case *Enemy:
		if t.blockedPart == bodyPart {
			damage *= 0.5
			fmt.Printf("%s заблокировал удар в %s! Урон уменьшен.\n", t.Name, bodyPart)
		}
		t.HP -= damage
		fmt.Printf("Вы нанесли %.1f урона врагу в %s (HP врага: %.1f)\n", damage, bodyPart, t.HP)
	case *Player:
		if t.blockedPart == bodyPart {
			damage *= 0.5
			fmt.Printf("%s заблокировал удар в %s! Урон уменьшен.\n", t.Name, bodyPart)
		}
		t.HP -= damage
		fmt.Printf("Враг нанес вам %.1f урона в %s (Ваше HP: %.1f)\n", damage, bodyPart, t.HP)
	}
}

func (p *Player) Block(bodyPart string) {
	p.blockedPart = bodyPart
	fmt.Printf("Вы блокируете %s\n", bodyPart)
}

func (e *Enemy) Hit(target Character, bodyPart string) {
	damage := float32(e.Strength)
	switch bodyPart {
	case "голова":
		damage *= 2
	case "тело":
		damage *= 1
	case "ноги":
		damage *= 1.5
	}

	switch t := target.(type) {
	case *Enemy:
		if t.blockedPart == bodyPart {
			damage *= 0.5
			fmt.Printf("%s заблокировал удар в %s! Урон уменьшен.\n", t.Name, bodyPart)
		}
		t.HP -= damage
		fmt.Printf("Враг нанес врагу %.1f урона в %s (HP врага: %.1f)\n", damage, bodyPart, t.HP)
	case *Player:
		if t.blockedPart == bodyPart {
			damage *= 0.5
			fmt.Printf("%s заблокировал удар в %s! Урон уменьшен.\n", t.Name, bodyPart)
		}
		t.HP -= damage
		fmt.Printf("Враг нанес вам %.1f урона в %s (Ваше HP: %.1f)\n", damage, bodyPart, t.HP)
	}
}

func (e *Enemy) Block(bodyPart string) {
	e.blockedPart = bodyPart
	fmt.Printf("Противник блокирует %s\n", bodyPart)
}

func randomBodyPart() string {
	parts := []string{"голова", "тело", "ноги"}
	return parts[rand.Intn(len(parts))]
}

func main() {
	player := Player{
		Name:     "Тето",
		HP:       10,
		Strength: 3,
	}

	enemy := Enemy{
		Name:     "Груша",
		HP:       10,
		Strength: 3,
	}

	var bodypartHit string
	var bodypartBlock string
	for player.HP > 0 && enemy.HP > 0 {
		fmt.Printf("\n=== Ход ===\n")
		fmt.Printf("Здоровье %s: %.1f\n", player.Name, player.HP)
		fmt.Printf("Здоровье %s: %.1f\n", enemy.Name, enemy.HP)

		fmt.Print("Куда атакуете противника? (голова/тело/ноги): ")
		fmt.Scanln(&bodypartHit)
		fmt.Print("Какую часть тела блокировать? (голова/тело/ноги): ")
		fmt.Scanln(&bodypartBlock)

		player.Block(bodypartBlock)
		player.Hit(&enemy, bodypartHit)

		if enemy.HP <= 0 {
			break
		}

		enemyHit := randomBodyPart()
		enemyBlock := randomBodyPart()
		enemy.Block(enemyBlock)
		enemy.Hit(&player, enemyHit)

		if player.HP <= 0 {
			break
		}
	}

	fmt.Printf("\n=== Бой окончен ===\n")
	fmt.Printf("Здоровье %s: %.1f\n", player.Name, player.HP)
	fmt.Printf("Здоровье %s: %.1f\n", enemy.Name, enemy.HP)

	if player.HP <= 0 && enemy.HP <= 0 {
		fmt.Println("Ничья! Оба умерли.")
	} else if player.HP <= 0 {
		fmt.Printf("Победитель: %s\n", enemy.Name)
		fmt.Printf("Проигравший: %s\n", player.Name)
		fmt.Println("Вы погибли")
	} else {
		fmt.Printf("Победитель: %s\n", player.Name)
		fmt.Printf("Проигравший: %s\n", enemy.Name)
	}

}
