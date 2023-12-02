package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	primesMiddle := primes[2:4]
	primesFront := primes[:3]
	primesBack := primes[3:]

	fmt.Println("primesMiddle: ", primesMiddle)
	fmt.Println("primesFront: ", primesFront)
	fmt.Println("primesBack: ", primesBack)

	slices := []int{2, 3, 5, 7, 11}
	fmt.Println("slices: ", slices)

	numbers := [][]int{{1, 2}, {3, 4}}
	fmt.Println("numbers: ", numbers)

	primesByMake := make([]int, 6)
	fmt.Println("primesByMake: ", primesByMake)
	primesByMake[0] = 2
	primesByMake[1] = 3
	fmt.Println("primesByMake: ", primesByMake)
}
