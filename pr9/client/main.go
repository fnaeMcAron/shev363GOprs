package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:8080")
	if err != nil {
		fmt.Println("ошибка подключения к серверу:", err)
		return
	}
	defer conn.Close()

	fmt.Println("подключено к серверу")

	messages := make(chan string, 5)

	go receiveMessages(conn, messages)
	go outputMessages(messages)

	// Чтение ввода пользователя
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("вы: ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		if msg != "" {
			fmt.Fprintln(conn, msg)
		}
	}
}

func receiveMessages(conn net.Conn, messages chan<- string) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("соединение с сервером разорвано")
			os.Exit(0)
		}

		messages <- strings.TrimSpace(msg)
	}
}

func outputMessages(messages <-chan string) {
	for msg := range messages {
		fmt.Println("\r" + msg)
		fmt.Print("вы: ")
	}
}
