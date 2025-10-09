package main

import (
	"errors"
	"fmt"
)

type Task struct {
	ID   int
	Text string
	Done bool
}

type TaskManager struct {
	tasks  map[int]*Task
	nextID int
}

func (tm *TaskManager) AddTask(text string) {
	if tm.tasks == nil {
		tm.tasks = make(map[int]*Task)
		tm.nextID = 1
	}

	task := &Task{
		ID:   tm.nextID,
		Text: text,
		Done: false,
	}

	tm.tasks[tm.nextID] = task
	tm.nextID = +1
	fmt.Println("+", text)
}

func (tm *TaskManager) RemoveTask(id int) error {
	task, ok := tm.tasks[id]
	if !ok {
		return errors.New("нет такой задачи")
	}

	delete(tm.tasks, id)
	fmt.Println("-", task.Text)
	return nil
}

func (tm *TaskManager) CompleteTask(id int) error {
	task, ok := tm.tasks[id]
	if !ok {
		return errors.New("нет такой задачи")
	}

	task.Done = true
	fmt.Println("Сделана задаа", task.Text)
	return nil
}

func (tm *TaskManager) ShowTasks() {
	for _, task := range tm.tasks {
		status := " "
		if task.Done {
			status = "ОК"
		}
		fmt.Printf("%s %d %s\n", status, task.ID, task.Text)
	}
}

func main() {
	manager := &TaskManager{}

	manager.AddTask("Купить хлеб")
	manager.AddTask("Сделать уроки")
	manager.AddTask("Сказать маме вызвать такси")

	manager.CompleteTask(1)
	manager.RemoveTask(2)

	manager.ShowTasks()
}
