package main

import "fmt"

type Stack struct {
	stack []int
}

func newStack() Stack {
	return Stack{stack: []int{}}
}

func (s *Stack) addToStack(item int) {
	s.stack = append(s.stack, item)
}

func (s *Stack) removeItemFromStack() (int, error) {
	if len(s.stack) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	temp := s.stack[len(s.stack)-1]
	s.stack = s.stack[0 : len(s.stack)-1]
	return temp, nil
}

func main() {
	aa := newStack()
	aa.addToStack(1)
	aa.addToStack(2)

	fmt.Println(aa.stack)
	fmt.Println(aa.removeItemFromStack())
	fmt.Println(aa.removeItemFromStack())
	fmt.Println(aa.removeItemFromStack())
}
