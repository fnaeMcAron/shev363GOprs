package server

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

var connected bool

func Start() {
	rand.Seed(time.Now().UnixNano())
	listener, _ := net.Listen("tcp", ":8080")
	fmt.Println("Сервер запускается...")

	for {
		conn, _ := listener.Accept()

		if connected {
			conn.Close()
			continue
		}

		connected = true
		go StartPvP(conn)
	}
}
