package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"io"
	"sync"
)

func CheckSumMD5(wg *sync.WaitGroup, jobs <-chan string) {
	 defer wg.Done()
    for filePath := range jobs {
        file, err := os.Open(filePath)
        if err != nil {
            fmt.Printf("Ошибка открытия файла %s: %v\n", filePath, err)
            continue
        }
        hasher := md5.New()
        if _, err := io.Copy(hasher, file); err != nil {
            fmt.Printf("Ошибка копирования данных файла %s: %v\n", filePath, err)
            file.Close()
            continue
        }
        file.Close()
        hashSum := hasher.Sum(nil)
        fmt.Printf("Файл: %s, MD5: %x\n", filePath, hashSum)
    }
}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan string, 5)
	Files := []string{
		"Files_For_Task5/log1.txt",
		"Files_For_Task5/log2.txt",
		"Files_For_Task5/log3.txt",
		"Files_For_Task5/log4.txt",
		"Files_For_Task5/log5.txt",
		"Files_For_Task5/log6.txt",
	}
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go CheckSumMD5(&wg, jobs)
	}
	for _, file := range Files {
		jobs <- file
	}
	close(jobs)
	wg.Wait()
}
