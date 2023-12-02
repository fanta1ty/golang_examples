package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main() {
	aa := Node{Value: 1}
	bb := Node{Value: 2}
	cc := Node{Value: 3}

	aa.Next = &bb
	bb.Next = &cc

	printNode(&aa)

	fmt.Println("===")
	nodeCount := lenNode(&aa)
	fmt.Println("Node count: ", nodeCount)

	fmt.Println("===")
	zz := Node{Value: 0}
	zz.Next = &aa
	printNode(&zz)

	fmt.Println("===")
	yy := Node{Value: 4}
	appendNode(&zz, &yy)
	printNode(&zz)

	fmt.Println("===")
	ww := Node{Value: 5}
	_ = insertNode(&zz, 1, &ww)
	printNode(&zz)

	fmt.Println("===")
	_ = deleteNode(&zz, 1)
	printNode(&zz)

}

func printNode(root *Node) {
	nodeWalk := root

	for nodeWalk.Next != nil {
		fmt.Println("node walk: ", nodeWalk.Value)
		nodeWalk = nodeWalk.Next
	}

	fmt.Println("node walk: ", nodeWalk.Value)
}

func lenNode(root *Node) int {
	nodeWalk := root
	count := 1

	for nodeWalk.Next != nil {
		count = count + 1
		nodeWalk = nodeWalk.Next
	}

	return count
}

func appendNode(root *Node, newNode *Node) {
	nodeWalk := root
	for nodeWalk.Next != nil {
		nodeWalk = nodeWalk.Next
	}
	nodeWalk.Next = newNode
}

func insertNode(root *Node, loc int, newNode *Node) error {
	nodeWalk := root
	counter := 0

	for nodeWalk.Next != nil {
		if counter == loc-1 {
			temp := nodeWalk.Next
			newNode.Next = temp
			nodeWalk.Next = newNode
			return nil
		}

		nodeWalk = nodeWalk.Next

		counter = counter + 1
	}

	return fmt.Errorf("Went pass no of element in list")
}

func deleteNode(root *Node, loc int) error {
	nodeWalk := root
	previousWalk := root
	counter := 0

	for nodeWalk.Next != nil {
		if counter == loc {
			previousWalk.Next = nodeWalk.Next
		}

		counter = counter + 1
		previousWalk = nodeWalk
		nodeWalk = nodeWalk.Next
	}

	return fmt.Errorf("Went pass the expected list")
}

func searchNode(root *Node, val int) *Node {
	nodeWalk := root

	for nodeWalk.Next != nil {
		if nodeWalk.Value == val {
			return nodeWalk
		}
		nodeWalk = nodeWalk.Next
	}

	return nil
}
