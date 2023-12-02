package main

import "fmt"

func bubbleSort(items []int) {
	if len(items) <= 1 {
		return
	}

	for {
		sortHappened := false

		for i := 0; i < len(items)-1; i++ {
			if items[i] > items[i+1] {
				temp := items[i]
				items[i] = items[i+1]
				items[i+1] = temp
				sortHappened = true
			}
		}

		if sortHappened == false {
			break
		}
	}
}

func mergeSort(items []int) []int {
	if len(items) <= 1 {
		return items
	}

	leftSide := mergeSort(items[0 : len(items)/2])
	rigtSide := mergeSort(items[len(items)/2 : len(items)])
	i := 0
	j := 0
	combined := []int{}

	for i < len(leftSide) || j < len(rigtSide) {
		if i >= len(leftSide) {
			combined = append(combined, rigtSide[j:]...)
			j = len(rigtSide)
			continue
		}

		if j >= len(rigtSide) {
			combined = append(combined, leftSide[i:]...)
			i = len(leftSide)
			continue
		}

		if leftSide[i] < rigtSide[i] {
			combined = append(combined, leftSide[i])
			i = i + 1
			continue
		}

		combined = append(combined, rigtSide[j])
		j = j + 1
	}

	return combined
}

func quickSort(values []int) []int {
	if len(values) <= 1 {
		return values
	}

	leftSide := []int{}
	rightSide := []int{}
	pivot := values[len(values)-1]

	for _, v := range values[0 : len(values)-1] {
		if v < pivot {
			leftSide = append(leftSide, v)
			continue
		}

		rightSide = append(rightSide, v)
	}

	sortedLeftSide := quickSort(leftSide)
	sortedRigtSide := quickSort(rightSide)

	sorted := append(sortedLeftSide, pivot)
	sorted = append(sorted, sortedRigtSide...)
	return sorted
}

func binarySearch(finding int, values []int) bool {
	if len(values) == 0 {
		return false
	}

	if len(values) == 1 {
		if values[0] == finding {
			return true
		}

		return false
	}

	found := false
	leftHalf := values[0 : len(values)/2]
	rightHalf := values[len(values)/2 : len(values)]

	if finding >= rightHalf[0] {
		found = binarySearch(finding, rightHalf)
	} else {
		found = binarySearch(finding, leftHalf)
	}

	return found
}

var store = map[int]int{1: 1, 2: 1}

func fibonacci(n int) int {
	if n <= 0 {
		return 0
	}

	if store[n] != 0 {
		return store[n]
	}

	val := fibonacci(n-1) + fibonacci(n-2)
	store[n] = val
	return val
}

func fibonacciTabulate(n int) int {
	if n <= 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	previous1 := 1
	previous2 := 1
	currentVal := 0

	for i := 3; i <= n; i++ {
		currentVal = previous1 + previous2
		previous1 = previous2
		previous2 = currentVal
	}
	return currentVal
}

func main() {
	values := []int{5, 3, 4, 1, 2}
	//fmt.Println(values)
	//bubbleSort(values)
	//fmt.Println(values)
	//sorted := mergeSort(values)
	//fmt.Println(sorted)
	sorted := quickSort(values)
	fmt.Println(sorted)

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	found := binarySearch(8, items)
	fmt.Printf("Was 8 found: %v\n", found)

	fmt.Println(fibonacci(100))
}
