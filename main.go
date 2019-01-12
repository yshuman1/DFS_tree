package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	var N int
	fmt.Scanf("%d", &N)
	fmt.Println(N)
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
