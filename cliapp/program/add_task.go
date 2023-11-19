package program

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/floire26/task-scheduler/dto"
	"github.com/floire26/task-scheduler/shared"
)

func AddTask(cfg *shared.Config) {
	var (
		name    string
		minute  string
		hour    string
		dom     string
		month   string
		dow     string
		funcOpt uint8
	)

	fmt.Printf("Input task name: ")
	fmt.Scan(&name)

	fmt.Printf("Input minute interval: ")
	fmt.Scan(&minute)

	fmt.Printf("Input hour interval: ")
	fmt.Scan(&hour)

	fmt.Printf("Input date of month interval: ")
	fmt.Scan(&dom)

	fmt.Printf("Input month interval: ")
	fmt.Scan(&month)

	fmt.Printf("Input week interval: ")
	fmt.Scan(&dow)

	fmt.Printf("Input function option: ")
	fmt.Scan(&funcOpt)

	data := dto.AddTaskRequest{
		Name:    name,
		Minute:  minute,
		Hour:    hour,
		Dom:     dom,
		Month:   month,
		Dow:     dow,
		FuncOpt: funcOpt,
	}

	reqBody, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	resBody := bytes.NewBuffer(reqBody)
	url := fmt.Sprintf("http://%s/tasks", cfg.Port)
	response, err := http.Post(url, "aplication/json", resBody)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("task addition to scheduler failed!")
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var out map[string]interface{}

	err = json.Unmarshal(responseData, &out)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(out["message"])
	fmt.Println("")
}
