package main

import (
    "fmt"
    "sync"
    "time"
)

type Task struct {
    id   int
    name string
}

func worker(id int, tasks <-chan Task, wg *sync.WaitGroup, results *sync.Map) {
    defer wg.Done()
    
    for task := range tasks {
        fmt.Printf("Worker %d начал: %s\n", id, task.name)
        time.Sleep(500 * time.Millisecond)
        fmt.Printf("Worker %d закончил: %s\n", id, task.name)
        
        
        results.Store(task.id, fmt.Sprintf("обработано worker %d", id))
    }
}

func producer(id int, tasks chan<- Task, numTasks int) {
    for i := 1; i <= numTasks; i++ {
        task := Task{
            id:   id*100 + i,
            name: fmt.Sprintf("Task-%d-%d", id, i),
        }
        tasks <- task
        fmt.Printf("Producer %d отправил: %s\n", id, task.name)
        time.Sleep(200 * time.Millisecond)
    }
}

func main() {

    tasks := make(chan Task, 10)
    
    
    var results sync.Map
    
    var wg sync.WaitGroup
    

    numWorkers := 3
    for i := 1; i <= numWorkers; i++ {
        wg.Add(1)
        go worker(i, tasks, &wg, &results)
    }
    
    
    numProducers := 2
    for i := 1; i <= numProducers; i++ {
        go producer(i, tasks, 4)
    }
    
    
    time.Sleep(3 * time.Second)
    close(tasks) 

    wg.Wait()

    fmt.Println("\nРезультаты")
    results.Range(func(key, value interface{}) bool {
        fmt.Printf("Задача %d: %s\n", key, value)
        return true
    })
}
