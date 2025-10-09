package main

//исправить дедлок
import (
	"fmt"
	"sync"
	"time"
)

type Error struct {
	stage string
	value int
	err   string
}

type Pipeline struct {
	errors []Error
	mu     sync.Mutex
}

func (ppl *Pipeline) addError(stage string, value int, err string) {
	ppl.mu.Lock()
	defer ppl.mu.Unlock()
	ppl.errors = append(ppl.errors, Error{stage, value, err})
}

func (ppl *Pipeline) reportErrors() {
	ppl.mu.Lock()
	defer ppl.mu.Unlock()

	fmt.Println("\nОшибки:")
	if len(ppl.errors) == 0 {
		fmt.Println("Пусто")
		return
	}

	for _, e := range ppl.errors {
		fmt.Printf("Стадия %s, значение %d: %s\n", e.stage, e.value, e.err)
	}
}

func Input(numbers <-chan int, results chan<- int, errors *Pipeline, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range numbers {
		if n%5 == 0 {
			errors.addError("Вход", n, "число делится на 5")
			continue
		}
		result := n + 5
		fmt.Printf("Вход (+5): %d -> %d\n", n, result)
		results <- result
		time.Sleep(100 * time.Millisecond)
	}
}

func Internal(numbers <-chan int, results chan<- int, errors *Pipeline, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range numbers {
		if n == 10 {
			errors.addError("Внутренний", n, "число равно 10")
			continue
		}
		result := n - 3
		fmt.Printf("Внутренний (-3): %d -> %d\n", n, result)
		results <- result
		time.Sleep(100 * time.Millisecond)
	}
}

func Output(numbers <-chan int, results chan<- int, errors *Pipeline, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range numbers {
		if n%2 == 0 {
			errors.addError("Выход", n, "число делится на 2")
			continue
		}
		result := n * 30
		fmt.Printf("Выход (*30): %d -> %d -> Конвейер закончил работу\n", n, result)
		results <- result
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	chan1 := make(chan int, 5)
	chan2 := make(chan int, 5)
	chan3 := make(chan int, 5)

	errors := &Pipeline{}
	var wg sync.WaitGroup

	wg.Add(3)
	go Input(chan1, chan2, errors, &wg)
	go Internal(chan2, chan3, errors, &wg)
	go Output(chan3, make(chan int), errors, &wg)

	for i := 1; i <= 15; i++ {
		chan1 <- i
	}
	close(chan1)

	wg.Wait()
	errors.reportErrors()
}
