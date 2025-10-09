package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, jobs <-chan string, once *sync.Once) {
	defer wg.Done()
	once.Do(func() {
		fmt.Println("Задача выполняется")
	})
}

func main() {
	var wg sync.WaitGroup
	var once sync.Once
	jobs := make(chan string, 5)
	tasks := []string{
		"задача 1",
		"задача 2",
		"задача 3",
		"задача 4",
		"задача 5",
	}

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(&wg, jobs, &once)
	}
	for _, task := range tasks {
		jobs <- task
	}
	close(jobs)
	wg.Wait()

}
