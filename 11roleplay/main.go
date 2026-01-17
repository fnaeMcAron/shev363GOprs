package main

import (
	"fmt"
	"math/rand"
)

type Character interface {
	Hit(target Character, bodyPart string)
	Block(bodyPart string)
}

type Item struct {
	Type    string
	Attack  int
	Defence int
	PlusHP  float32
}

type Player struct {
	Name        string
	HP          float32
	Strength    int
	blockedPart string
	Inventory   []Item
	Equipment   []Item
}

type Enemy struct {
	Name        string
	HP          float32
	Strength    int
	blockedPart string
	Inventory   []Item
	Equipment   []Item
	Item        Item
}

func (p *Player) Hit(target Character, bodyPart string) {
	damage := float32(p.Strength)

	for _, item := range p.Equipment {
		if item.Type == "оружие" {
			damage += float32(item.Attack)
		}
	}

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
		enemyDefence := 0
		for _, item := range t.Equipment {
			if item.Type == "броня" {
				enemyDefence += item.Defence
			}
		}
		damage -= float32(enemyDefence)
		if damage < 0 {
			damage = 0
		}

		if t.blockedPart == bodyPart {
			damage *= 0.5
			fmt.Printf("%s заблокировал удар в %s! Урон уменьшен.\n", t.Name, bodyPart)
		}
		t.HP -= damage
		fmt.Printf("Вы нанесли %.1f урона врагу в %s (HP врага: %.1f)\n", damage, bodyPart, t.HP)
	case *Player:
		playerDefence := 0
		for _, item := range t.Equipment {
			if item.Type == "броня" {
				playerDefence += item.Defence
			}
		}
		damage -= float32(playerDefence)
		if damage < 0 {
			damage = 0
		}

		if t.blockedPart == bodyPart {
			damage *= 0.5
			fmt.Printf("%s заблокировал удар в %s! Урон уменьшен.\n", t.Name, bodyPart)
		}
		t.HP -= damage
		fmt.Printf("%s нанес %.1f урона %s в %s (HP %s: %.1f)\n", p.Name, damage, t.Name, bodyPart, t.Name, t.HP)
	}
}

func (p *Player) Block(bodyPart string) {
	p.blockedPart = bodyPart
	fmt.Printf("%s блокирует %s\n", p.Name, bodyPart)
}

func (e *Enemy) Hit(target Character, bodyPart string) {
	damage := float32(e.Strength)

	for _, item := range e.Equipment {
		if item.Type == "оружие" {
			damage += float32(item.Attack)
		}
	}

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
		playerDefence := 0
		for _, item := range t.Equipment {
			if item.Type == "броня" {
				playerDefence += item.Defence
			}
		}
		damage -= float32(playerDefence)
		if damage < 0 {
			damage = 0
		}

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

func (p *Player) TakeOff() {
	if len(p.Equipment) == 0 {
		fmt.Printf("На %s ничего не надето!\n", p.Name)
		return
	}

	fmt.Printf("=== Экипированные предметы %s ===\n", p.Name)
	for i, item := range p.Equipment {
		fmt.Printf("%d. %s", i+1, item.Type)
		if item.Type == "оружие" {
			fmt.Printf(" (Атака: +%d)", item.Attack)
		} else if item.Type == "броня" {
			fmt.Printf(" (Защита: +%d)", item.Defence)
		}
		fmt.Println()
	}

	fmt.Print("Введите номер предмета для снятия (0 для отмены): ")
	var choice int
	fmt.Scanln(&choice)

	if choice == 0 {
		return
	}

	if choice < 1 || choice > len(p.Equipment) {
		fmt.Println("Неверный выбор!")
		return
	}

	itemIndex := choice - 1
	item := p.Equipment[itemIndex]

	p.Inventory = append(p.Inventory, item)

	p.Equipment = append(p.Equipment[:itemIndex], p.Equipment[itemIndex+1:]...)

	fmt.Printf("%s снял %s\n", p.Name, item.Type)
}

func (p *Player) Equip() {
	if len(p.Inventory) == 0 {
		fmt.Printf("Инвентарь %s пуст!\n", p.Name)
		return
	}

	fmt.Printf("=== Инвентарь %s ===\n", p.Name)
	for i, item := range p.Inventory {
		fmt.Printf("%d. %s", i+1, item.Type)
		switch item.Type {
		case "оружие":
			fmt.Printf(" (Атака: +%d)", item.Attack)
		case "броня":
			fmt.Printf(" (Защита: +%d)", item.Defence)
		case "применяемый предмет":
			fmt.Printf(" (Восстановление HP: +%.1f)", item.PlusHP)
		}
		fmt.Println()
	}

	fmt.Print("Введите номер предмета для использования (0 для отмены): ")
	var choice int
	fmt.Scanln(&choice)

	if choice == 0 {
		return
	}

	if choice < 1 || choice > len(p.Inventory) {
		fmt.Println("Неверный выбор!")
		return
	}

	itemIndex := choice - 1
	item := p.Inventory[itemIndex]

	if item.Type == "применяемый предмет" {
		p.HP += item.PlusHP
		fmt.Printf("%s использовал %s и восстановил %.1f HP\n", p.Name, item.Type, item.PlusHP)
		fmt.Printf("HP %s: %.1f\n", p.Name, p.HP)

		p.Inventory = append(p.Inventory[:itemIndex], p.Inventory[itemIndex+1:]...)
		return
	}

	for _, equippedItem := range p.Equipment {
		if equippedItem.Type == item.Type {
			fmt.Printf("У %s уже надет %s! Сначала снимите его.\n", p.Name, item.Type)
			return
		}
	}

	p.Equipment = append(p.Equipment, item)

	p.Inventory = append(p.Inventory[:itemIndex], p.Inventory[itemIndex+1:]...)

	fmt.Printf("%s надел %s\n", p.Name, item.Type)
}

func generateRandomItem() Item {
	itemTypes := []string{"оружие", "броня", "применяемый предмет"}
	itemType := itemTypes[rand.Intn(len(itemTypes))]

	var item Item
	item.Type = itemType

	switch itemType {
	case "оружие":
		item.Attack = rand.Intn(5) + 1
	case "броня":
		item.Defence = rand.Intn(5) + 1
	case "применяемый предмет":
		item.PlusHP = float32(rand.Intn(10)+5) / 2
	}

	return item
}

func pvpMode() {
	pl1name := ""
	pl2name := ""
	fmt.Printf("Игрок 1, введите свое имя\n")
	fmt.Scanln(&pl1name)
	fmt.Printf("Игрок 2, введите свое имя\n")
	fmt.Scanln(&pl2name)
	weapons := []Item{
		{Type: "оружие", Attack: 3},
		{Type: "оружие", Attack: 5},
		{Type: "оружие", Attack: 2},
		{Type: "оружие", Attack: 4},
		{Type: "оружие", Attack: 6},
	}

	armors := []Item{
		{Type: "броня", Defence: 2},
		{Type: "броня", Defence: 4},
		{Type: "броня", Defence: 3},
		{Type: "броня", Defence: 5},
		{Type: "броня", Defence: 1},
	}

	consumables := []Item{
		{Type: "применяемый предмет", PlusHP: 5},
		{Type: "применяемый предмет", PlusHP: 3},
		{Type: "применяемый предмет", PlusHP: 7},
	}

	players := [2]Player{
		{
			Name:      pl1name,
			HP:        15,
			Strength:  3,
			Inventory: []Item{weapons[0], armors[0], consumables[0]},
		},
		{
			Name:      pl2name,
			HP:        15,
			Strength:  3,
			Inventory: []Item{weapons[1], armors[1], consumables[1]},
		},
	}

	var bodypartHit string
	var bodypartBlock string
	var action string
	currentPlayer := 0
	for players[0].HP > 0 && players[1].HP > 0 {
		fmt.Printf("\n=== Ход %s ===\n", players[currentPlayer].Name)
		fmt.Printf("Здоровье %s: %.1f\n", players[0].Name, players[0].HP)
		fmt.Printf("Здоровье %s: %.1f\n", players[1].Name, players[1].HP)

		fmt.Printf("%s, выберите действие (атака/экипировка/снять): ", players[currentPlayer].Name)
		fmt.Scanln(&action)

		switch action {
		case "экипировка":
			players[currentPlayer].Equip()
			continue
		case "снять":
			players[currentPlayer].TakeOff()
			continue
		case "атака":
			// игнорирование
		default:
			fmt.Println("Неизвестное действие, продолжаем атаку")
		}

		targetIndex := (currentPlayer + 1) % 2
		fmt.Printf("Куда атакуете %s? (голова/тело/ноги): ", players[targetIndex].Name)
		fmt.Scanln(&bodypartHit)
		fmt.Print("Какую часть тела блокировать? (голова/тело/ноги): ")
		fmt.Scanln(&bodypartBlock)

		players[currentPlayer].Block(bodypartBlock)
		players[currentPlayer].Hit(&players[targetIndex], bodypartHit)

		if players[targetIndex].HP <= 0 {
			break
		}

		currentPlayer = (currentPlayer + 1) % 2
	}

	fmt.Printf("\n=== Бой окончен ===\n")
	fmt.Printf("Здоровье %s: %.1f\n", players[0].Name, players[0].HP)
	fmt.Printf("Здоровье %s: %.1f\n", players[1].Name, players[1].HP)

	if players[0].HP <= 0 && players[1].HP <= 0 {
		fmt.Println("Ничья - Оба игрока умерли. Хах.")
	} else if players[0].HP <= 0 {
		fmt.Printf("Победитель: %s\n", players[1].Name)
		fmt.Printf("Проигравший: %s\n", players[0].Name)
	} else {
		fmt.Printf("Победитель: %s\n", players[0].Name)
		fmt.Printf("Проигравший: %s\n", players[1].Name)
	}
}

func pveMode() {
	weapons := []Item{
		{Type: "оружие", Attack: 3},
		{Type: "оружие", Attack: 5},
		{Type: "оружие", Attack: 2},
		{Type: "оружие", Attack: 4},
		{Type: "оружие", Attack: 6},
	}

	armors := []Item{
		{Type: "броня", Defence: 2},
		{Type: "броня", Defence: 4},
		{Type: "броня", Defence: 3},
		{Type: "броня", Defence: 5},
		{Type: "броня", Defence: 1},
	}

	consumables := []Item{
		{Type: "применяемый предмет", PlusHP: 5},
		{Type: "применяемый предмет", PlusHP: 3},
		{Type: "применяемый предмет", PlusHP: 7},
	}

	player := Player{
		Name:      "Тето",
		HP:        10,
		Strength:  3,
		Inventory: []Item{weapons[0], armors[0], consumables[0]},
	}

	enemy := Enemy{
		Name:     "Груша",
		HP:       10,
		Strength: 3,
		Item:     generateRandomItem(),
	}

	var bodypartHit string
	var bodypartBlock string
	var action string

	for player.HP > 0 && enemy.HP > 0 {
		fmt.Printf("\n=== Ход ===\n")
		fmt.Printf("Здоровье %s: %.1f\n", player.Name, player.HP)
		fmt.Printf("Здоровье %s: %.1f\n", enemy.Name, enemy.HP)

		fmt.Print("Выберите действие (атака/экипировка/снять): ")
		fmt.Scanln(&action)

		switch action {
		case "экипировка":
			player.Equip()
			continue
		case "снять":
			player.TakeOff()
			continue
		case "атака":
			// игнорирование 2
		default:
			fmt.Println("Неизвестное действие, продолжаем атаку")
		}

		fmt.Print("Куда атакуете противника? (голова/тело/ноги): ")
		fmt.Scanln(&bodypartHit)
		fmt.Print("Какую часть тела блокировать? (голова/тело/ноги): ")
		fmt.Scanln(&bodypartBlock)

		player.Block(bodypartBlock)
		player.Hit(&enemy, bodypartHit)

		if enemy.HP <= 0 {
			if enemy.Item.Type != "" {
				player.Inventory = append(player.Inventory, enemy.Item)
				fmt.Printf("Вы получили трофей: %s!\n", enemy.Item.Type)
			}
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

func main() {
	rand.Seed(42)

	var mode string
	fmt.Println("Выберите режим игры:")
	fmt.Println("1 - PvP (игрок против игрока)")
	fmt.Println("2 - PvE (игрок против компьютера)")
	fmt.Print("Ваш выбор: ")
	fmt.Scanln(&mode)

	switch mode {
	case "1", "pvp", "PvP":
		fmt.Println("\n=== Режим PvP ===")
		pvpMode()
	case "2", "pve", "PvE":
		fmt.Println("\n=== Режим PvE ===")
		pveMode()
	default:
		fmt.Println("Неизвестный режим, запускается PvE")
		pveMode()
	}
}
