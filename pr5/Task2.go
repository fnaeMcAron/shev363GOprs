package main

import(
	"time"
	"fmt"
	"os"
) 

func main(){
	done := make(chan struct{})

	go func(){
		for{
			time.Sleep(1 * time.Second)
			fmt.Println("Сообщение")
		}
	}()
	go func(){
		os.Stdin.Read(make([]byte, 1))
		done <- struct{}{}
	}()
	select {
	case <-done: 
			fmt.Println("Закрытие")
	}
}