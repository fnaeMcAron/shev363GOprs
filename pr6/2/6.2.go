package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex
var cache = map[int]string{
	0: "Алиса",
	1: "Боб",
	2: "ЕВГЕНИЙ",
}

func Reader(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	mu.RLock()
	defer mu.RUnlock()
	time.Sleep(2 * time.Second)
	fmt.Println(cache[id])
	fmt.Printf("Кэш прочитан воркером с номером %d \n", id)
}

func Writer(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	mu.Lock()
	time.Sleep(2 * time.Second)
	cache[id] = "Боб"
	fmt.Println("Перезаписанный кэш:" + cache[id])
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	Worker := 3
	for i := 0; i <= Worker; i++ {
		wg.Add(1)
		go Reader(&wg, i)
	}
	go Writer(&wg, 1)
	wg.Wait()
}
