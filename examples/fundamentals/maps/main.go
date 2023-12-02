package main

import "fmt"

func main() {
	mappedItems := map[string]string{}
	mappedItems["testKey"] = "testValue"
	fmt.Println("mappedItems: ", mappedItems)

	mappedItemsByMake := make(map[string]string)
	mappedItemsByMake["testKey"] = "testValue"
	fmt.Println("mappedItemsByMake: ", mappedItemsByMake)

	//mappedItems := make(map[string]int)
	//mappedItems := make(map[int]string)
	//mappedItems := make(map[string]bool)
	//mappedItems := make(map[string]func(a int) int)
}
