package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	test := &Node{
		Value: 123,
	}
	right := &Node{
		Value: 333,
	}
	test.Right = right
	fmt.Println(test)
	printNode(test)
}
func printNode(n *Node) {
	fmt.Print("Value: ", n.Value)
	if n.Left != nil {
		fmt.Print(" Left: ", n.Left.Value)
	}
	if n.Right != nil {
		fmt.Print(" Right: ", n.Right.Value)
	}
	fmt.Println()
}
