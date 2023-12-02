package main

import (
	"errors"
	"fmt"
)

func main() {

}

func exampleFunc() (int, error) {
	return 1, errors.New("This is an example error")
}

func exampleFunc2() (int, error) {
	return 1, fmt.Errorf("This is an example error")
}
