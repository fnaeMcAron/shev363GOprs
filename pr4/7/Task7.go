package main

import (
	"fmt"
)

type Settings struct {
	Theme         string
	Language      string
	Notifications bool
}

func Manager(setCh <-chan Settings, getCh chan<- Settings) {
	current := Settings{
		Theme:         "Темная",
		Language:      "en",
		Notifications: true,
	}

	for {
		select {
		case newSettings := <-setCh:
			current = newSettings
		case getCh <- current:
		}
	}
}

func setSet(ch chan<- Settings, settings Settings) {
	ch <- settings
}

func getSet(ch <-chan Settings) Settings {
	return <-ch
}

func main() {
	setCh := make(chan Settings)
	getCh := make(chan Settings)

	go Manager(setCh, getCh)

	fmt.Println("Заводские настройки:", getSet(getCh))

	newSet := Settings{
		Theme:         "Светлая",
		Language:      "ru",
		Notifications: false,
	}
	setSet(setCh, newSet)

	fmt.Println("Обновленные настройки:", getSet(getCh))
}
