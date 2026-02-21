package server

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"rpg-server/game"
)

type Turn struct {
	Hit   string `json:"hit"`
	Block string `json:"block"`
}

func StartPvP(conn net.Conn) {
	defer conn.Close()

	serverPlayer := &game.Player{
		Name:     "Го",
		HP:       100,
		Strength: 15,
	}

	clientPlayer := &game.Player{
		Name:     "Юнити",
		HP:       100,
		Strength: 12,
	}

	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	incoming := make(chan Message)

	go func() {
		for {
			var msg Message
			err := decoder.Decode(&msg)
			if err != nil {
				//fmt.Printf("Ошибка декодирования: %v\n", err)
				close(incoming)
				return
			}
			//fmt.Printf("ПОЛУЧЕНО: type=%s, data=%v (тип: %T)\n", msg.Type, msg.Data, msg.Data)
			incoming <- msg
		}
	}()

	firstTurnServer := rand.Intn(2) == 0

	for serverPlayer.IsAlive() && clientPlayer.IsAlive() {

		fmt.Println("Ожидание хода клиента...")

		var clientTurn Turn

		for {
			msg := <-incoming

			if msg.Type == "chat" {
				fmt.Println("Клиент:", msg.Data)
				continue
			}

			if msg.Type == "turn" {
				data, ok := msg.Data.(map[string]interface{})
				if !ok {
					fmt.Println("Ошибка: data не является map")
					continue
				}
				clientTurn.Hit = data["hit"].(string)
				clientTurn.Block = data["block"].(string)
				break
			}
		}

		fmt.Println("Куда бить (голова/тело/ноги):")
		fmt.Scan(&serverPlayer.HitPart)

		fmt.Println("Что блокировать (голова/тело/ноги):")
		var block string
		fmt.Scan(&block)
		serverPlayer.Block(block)

		clientPlayer.HitPart = clientTurn.Hit
		clientPlayer.Block(clientTurn.Block)

		if firstTurnServer {

			serverPlayer.Hit(clientPlayer)

			if clientPlayer.IsAlive() {
				clientPlayer.Hit(serverPlayer)
			}

		} else {

			clientPlayer.Hit(serverPlayer)

			if serverPlayer.IsAlive() {
				serverPlayer.Hit(clientPlayer)
			}
		}

		firstTurnServer = !firstTurnServer

		state := map[string]interface{}{
			"serverHP": serverPlayer.HP,
			"clientHP": clientPlayer.HP,
		}

		encoder.Encode(Message{
			Type: "state",
			Data: state,
		})
	}

	fmt.Println("le fin.")
}
