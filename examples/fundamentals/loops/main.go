package main

import "fmt"

type Subtask struct {
	Param1 string
	Param2 string
	Status string
}

func main() {
	allTasks := []Subtask{{Status: "incomplete"}, {Status: "completed"}}

	for idx, x := range allTasks {
		println("index: ", idx)
		if x.Status != "completed" {
			fmt.Println("The main task is sill incomplete")
		}
	}

	items := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	for idx, x := range items {
		fmt.Println("Idx: ", idx)
		fmt.Print(x)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
