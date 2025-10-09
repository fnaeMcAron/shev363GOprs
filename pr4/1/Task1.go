package main

import (
    "fmt"
    "sync"
    "time"
)

func PrintNumbers(number int, wg *sync.WaitGroup){
	defer wg.Done()
    time.Sleep(1 * time.Second)
    fmt.Println(number)
}


func main(){
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
        go PrintNumbers(i, &wg)
		wg.Wait()
	}

}