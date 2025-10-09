package main

import (
    "fmt"
    "sync"
)

func worker(wg *sync.WaitGroup, jobs <-chan int, results chan<- int) {
	defer wg.Done()
    for num := range jobs {
        results <- num * num
    }
}

func main(){
	var wg sync.WaitGroup
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	for w := 1; w <= 3; w++ {
		wg.Add(1)
    	go worker(&wg, jobs, results)
    }
	for j := 1; j <= 10; j++ {
        jobs <- j
    }
    close(jobs)
	wg.Wait()
	close(results)
	for number := range results{
        fmt.Println(number)
    }
}