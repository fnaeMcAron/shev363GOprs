package main

import(
	"fmt"
	"time"
) 

func main(){
	highPriority := make(chan string, 1)
	lowPriority := make(chan string, 1)

	go func(){
		time.Sleep(1 * time.Second)
		highPriority <- "Задача с высоким приоритетом"
	}()
	go func(){
		time.Sleep(2 * time.Second)
		lowPriority <- "Задача с низким приоритетом"
	}()

	for i :=0; i < 2; i++{
		select {
		case msg := <-highPriority:
			fmt.Println(msg)
		case msg := <-lowPriority:
			fmt.Println(msg)
		}
	}
}