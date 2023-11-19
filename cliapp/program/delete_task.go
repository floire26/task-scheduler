package program

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/floire26/task-scheduler/shared"
)

func DeleteTask(cfg *shared.Config) {
	var id int

	fmt.Printf("Input task ID: ")
	fmt.Scan(&id)
	fmt.Println("")

	url := fmt.Sprintf("http://%s/tasks/%d", cfg.Port, id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("task deletion from scheduler failed!")
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
