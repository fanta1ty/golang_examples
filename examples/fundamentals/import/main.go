package main

import (
	"fmt"
	temp "github.com/go-yaml/yaml"
	"math/rand"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(rand.Int())

	output, _ := temp.Marshal(map[string]string{"a": "b", "c": "d"})
	fmt.Println(string(output))
}
