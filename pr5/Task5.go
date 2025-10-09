package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(wg *sync.WaitGroup, done <-chan struct{}, Request <-chan string, NumberWorker int) {
    defer wg.Done()
    for {
        select {
        case answer, ok := <-Request:
            if !ok {
                fmt.Println("Канал запросов закрыт, ожидание завершения работы")
                return
            }
			fmt.Printf("№%d ", NumberWorker)
            fmt.Println("Воркер", answer)
        case <-done:
			fmt.Printf("№%d ", NumberWorker)
            fmt.Println("Воркер: все запросы выполнены, отправлен сигнал завершения")
            return
        default:
            time.Sleep(1 * time.Second)
            fmt.Println("Ожидание запроса")
        }
    }
}

func main() {
    var wg sync.WaitGroup
    Request := make(chan string)
    done := make(chan struct{})
    numWorkers := 3
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go worker(&wg, done, Request, i)
    }

    go func() {
        for i := 0; i < 10; i++ {
            Request <- fmt.Sprintf("Запрос %d получен", i+1)
            time.Sleep(500 * time.Millisecond) 
        }
        close(Request)
    }()


    time.Sleep(3 * time.Second)
    wg.Wait()
    close(done)
    fmt.Println("Завершение работы сервера")
}