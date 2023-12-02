package main

import "fmt"

func main() {
	exampleFunc()
	anotherExampleFunc()
	exampleFunc2("Hello", 10)
	exampleFunc3(true, "Hello", "World", 1, 2)
	exampleFunc4("h", "e", "l", "l", "o", "!")

	stringList := []string{"aa", "bb", "cc"}
	exampleFunc4(stringList...)

	exampleFunc5("Hello", subFunc)

	result := exampleFunc6("Hello World")
	fmt.Println("result: ", result)

	result2, index := exampleFunc7("Hello World")
	fmt.Println("result2: ", result2, "index: ", index)
}

func exampleFunc() {
	fmt.Println("exampleFunc")
}

func anotherExampleFunc() {
	fmt.Println("anotherExampleFunc")
	a := 10
	if a == 10 {
		fmt.Println("Hello from func")
		return
	}
	fmt.Println("Bye")
}

func exampleFunc2(a string, b int) {
	fmt.Println(a)
	fmt.Println(b)
}

func exampleFunc3(a bool, b, c string, d, e int) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

func exampleFunc4(a ...string) {
	fmt.Println(a)
}

func subFunc(a string) {
	fmt.Println("subFunc" + " " + a)
}

func exampleFunc5(a string, b func(c string)) {
	fmt.Println(a)
	b(a)
}

func exampleFunc6(s string) string {
	return s
}

func exampleFunc7(s string) (string, int) {
	return s, 1
}
