package main

import(
	"fmt"
) 

func main(){
	data := make(chan string, 1)
	select {
	case info := <-data: 
		if info != "" {
			fmt.Println("Данные есть")
		}
	default:
		fmt.Println("Данных нет")
	}
}