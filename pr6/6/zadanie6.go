package main

import (
    "fmt"
    "sync"
    "time"
)

type Logger struct {
    mu sync.Mutex
}

func (l *Logger) Log(workerID int, message string) {
    l.mu.Lock()
    defer l.mu.Unlock()
    
    fmt.Printf("[%s] Worker %d: %s\n", time.Now().Format("15:04:05"), workerID, message)
}

func worker(id int, logger *Logger, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for i := 1; i <= 5; i++ {
        logger.Log(id, fmt.Sprintf("Message %d", i))
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    logger := &Logger{}
    var wg sync.WaitGroup
    

    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, logger, &wg)
    }
    
    wg.Wait()
    fmt.Println("все логи записаны")
}
