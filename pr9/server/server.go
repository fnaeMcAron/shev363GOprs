package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	clients     = make(map[net.Conn]bool)
	clientsLock sync.RWMutex
	messages    = make(chan string, 5)
)

func main() {
	listener, err := net.Listen("tcp4", ":8080")
	if err != nil {
		fmt.Println("ошибка запуска сервера:", err)
		return
	}
	defer listener.Close()

	go outputMessages()
	go serverInput()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("ошибка подключения:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	clientsLock.Lock()
	clients[conn] = true
	clientsLock.Unlock()

	messages <- fmt.Sprintf("\nклиент подключился: %s", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		msg = strings.TrimSpace(msg)
		if msg == "" {
			continue
		}

		fullMsg := fmt.Sprintf("[%s]: %s", conn.RemoteAddr(), msg)
		messages <- fullMsg

		broadcast(fullMsg, conn)
	}

	clientsLock.Lock()
	delete(clients, conn)
	clientsLock.Unlock()

	messages <- fmt.Sprintf("\nклиент отключился: %s", conn.RemoteAddr())
}

func broadcast(msg string, sender net.Conn) {
	clientsLock.RLock()
	defer clientsLock.RUnlock()

	for client := range clients {
		fmt.Fprintln(client, msg)
	}
}

func outputMessages() {
	for msg := range messages {
		fmt.Println(msg)
	}
}

func serverInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		if msg != "" {
			fullMsg := fmt.Sprintf("\n[ADMINISTRATOR 😈]: %s", msg)
			messages <- fullMsg
			broadcast(fullMsg, nil)
		}
	}
}
