package main

import (
	"fmt"
	"sync"
	"time"
)

type Metrics struct {
	mu                                     sync.RWMutex
	errorCount, successCount, requestCount int
	totalTime                              time.Duration
}

func (metr *Metrics) AddSuccess(duration time.Duration) {
	metr.mu.Lock()
	defer metr.mu.Unlock()
	metr.successCount++
	metr.requestCount++
	metr.totalTime += duration
}

func (metr *Metrics) AddError() {
	metr.mu.Lock()
	defer metr.mu.Unlock()
	metr.errorCount++
	metr.requestCount++
}

func (metr *Metrics) Report() {
	metr.mu.RLock()
	defer metr.mu.RUnlock()

	avgTime := time.Duration(0)
	if metr.requestCount > 0 {
		avgTime = metr.totalTime / time.Duration(metr.requestCount)
	}

	fmt.Println("\nМетрики")
	fmt.Printf("Успешных запросов: %d\n", metr.successCount)
	fmt.Printf("Ошибок: %d\n", metr.errorCount)
	fmt.Printf("Всего запросов: %d\n", metr.requestCount)
	fmt.Printf("Среднее время: %v\n", avgTime)

	successRate := 0.0
	if metr.requestCount > 0 {
		successRate = float64(metr.successCount) / float64(metr.requestCount) * 100
	}
	fmt.Printf("Процент успешных: %.1f%%\n", successRate)
}

func makeRequest(workerID int, metrics *Metrics, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		start := time.Now()

		time.Sleep(time.Duration(100+workerID*10) * time.Millisecond)

		if i%3 != 0 {
			metrics.AddSuccess(time.Since(start))
			fmt.Printf("%d: Запрос №%d УСПЕШНО\n", workerID, i)
		} else {
			metrics.AddError()
			fmt.Printf("%d: Запрос №%d НЕУДАЧА :D\n", workerID, i)
		}
	}
}

func main() {
	metrics := &Metrics{}
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go makeRequest(i, metrics, &wg)
	}

	wg.Wait()
	metrics.Report()
}
