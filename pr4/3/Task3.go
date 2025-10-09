package main

import (
    "fmt"
	"sync"
	"time"
)

func ApplyRequest(tick <-chan time.Time, wg *sync.WaitGroup, id int){
	defer wg.Done()
	<- tick
	fmt.Printf("Запрос принят - %d \n", id)
}

func main(){
	var wg sync.WaitGroup
	tick := time.Tick(200*time.Millisecond)
	Requests := 15

	for i := 1; i <= Requests; i++{
		wg.Add(1)
		ApplyRequest(tick, &wg, i)	
	}
}