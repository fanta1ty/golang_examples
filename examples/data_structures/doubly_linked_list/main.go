package main

import "fmt"

type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

func main() {
	aa := Node{Value: 1}
	bb := Node{Value: 2}
	cc := Node{Value: 3}

	aa.Next = &bb
	bb.Next = &cc
	bb.Previous = &aa
	cc.Previous = &bb

	zz := Node{Value: 4}
	_ = insertNode(&aa, &zz, 3)
	printNode(&aa)
}

func printNode(root *Node) {
	nodeWalk := root

	for nodeWalk.Next != nil {
		fmt.Println("node walk: ", nodeWalk.Value)
		nodeWalk = nodeWalk.Next
	}

	fmt.Println("node walk: ", nodeWalk.Value)
}

func insertNode(root *Node, newNode *Node, loc int) *Node {
	counter := 0
	n := root
	var p *Node
	for n != nil {
		if counter == loc {
			newNode.Next = n
			newNode.Previous = p

			if n.Next != nil {
				n.Next.Previous = newNode
			}

			if p != nil {
				p.Next = newNode
				return root
			}

			return newNode
		}

		p = n
		n = n.Next
		counter = counter + 1
	}

	if counter == loc {
		newNode.Previous = p
		p.Next = newNode
		return root
	}

	return nil
}

func deleteNode(root *Node, loc int) *Node {
	counter := 0
	n := root
	var p *Node

	for n != nil {
		if counter == loc {
			if p == nil {
				temp := n.Next
				n.Next = nil
				n.Previous = nil
				return temp
			}
			p.Next = n.Next
			n.Next = nil
			n.Previous = nil
			return root
		}

		p = n
		n = n.Next
		counter = counter + 1
	}

	return nil
}
