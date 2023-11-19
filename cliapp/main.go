package main

import (
	"fmt"

	"github.com/floire26/task-scheduler/cliapp/program"
	"github.com/floire26/task-scheduler/shared"
)

func main() {
	cfg := shared.LoadConfig("../.env")
	fmt.Println("Task Scheduler CLI")
	opt := 1
	for opt != 0 {
		fmt.Println("Input an option: ")
		fmt.Println("1. Add a task")
		fmt.Println("2. Delete a task")
		fmt.Println("0. Exit")
		fmt.Printf("Your input: ")
		fmt.Scan(&opt)
		switch opt {
		case 1:
			program.AddTask(cfg)
		case 2:
			program.DeleteTask(cfg)
		}
	}
}
