package main

import (
	"fmt"
	"os"
	"time"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
}

// СТРУКТУРЫ
type FileLogger struct {
	filename string
	file     *os.File
}

type ConsoleLogger struct {
	name string
}

// 1000 и 1 метод
func (cl *ConsoleLogger) Info(msg string) {
	fmt.Printf("[%s] [INFO] [%s] %s\n", time.Now(), cl.name, msg)
}

func (cl *ConsoleLogger) Error(msg string) {
	fmt.Printf("[%s] [ERROR] [%s] %s\n", time.Now(), cl.name, msg)
}

func (cl *ConsoleLogger) Debug(msg string) {
	fmt.Printf("[%s] [DEBUG] [%s] %s\n", time.Now(), cl.name, msg)
}

func (fl *FileLogger) Info(msg string) {
	log := fmt.Sprintf("[%s] [INFO] %s\n", time.Now(), msg)
	fl.file.WriteString(log)
}

func (fl *FileLogger) Error(msg string) {
	log := fmt.Sprintf("[%s] [ERROR] %s\n", time.Now(), msg)
	fl.file.WriteString(log)
}

func (fl *FileLogger) Debug(msg string) {
	log := fmt.Sprintf("[%s] [DEBUG] %s\n", time.Now(), msg)
	fl.file.WriteString(log)
}

func (fl *FileLogger) Close() error {
	return fl.file.Close()
}

// МЕЕЙН
func main() {
	//что такое флаг в OpenFile() я так и не понял, поэтому Create()
	file1, err := os.Create("f1")
	file2, err := os.Create("f2")

	cl := ConsoleLogger{name: "АБв"}
	cl2 := ConsoleLogger{name: "ффыв"}
	cl.Debug("test")
	cl2.Error("ertest")
	cl.Info("важная информация")
	fl1 := FileLogger{filename: "тест1", file: file1}
	fl2 := FileLogger{filename: "тест2", file: file2}
	fl1.Debug("adasd")
	fl2.Error("adasawqw'")
	fmt.Print(err)
	fmt.Println(" 😞")
}
