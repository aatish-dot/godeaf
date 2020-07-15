package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func main() {
	nodes := readNodes()
	for _, v := range nodes {
		printnode(v)
	}

}
func readNodes() []Node {
	var N int
	fmt.Scanf("%d", &N)
	fmt.Println("N is ", N)
	var nodes []Node = make([]Node, N)
	for i := 0; i < N; i++ {
		var val, indexLeft, indexRight int
		fmt.Scanf("%d %d %d", &val, &indexLeft, &indexRight)
		fmt.Println(val, indexLeft, indexRight)
		nodes[i].value = val
		if indexLeft >= 0 {
			nodes[i].left = &nodes[indexLeft]

		}

		if indexRight >= 0 {
			nodes[i].right = &nodes[indexRight]

		}

	}

	return nodes
}

func printnode(N Node) {
	fmt.Print("Value of Node is ", N.value)
	if N.left != nil {
		fmt.Print("Left of Node is ", N.left.value)
	}
	if N.right != nil {
		fmt.Print("Right of Node is ", N.right.value)
	}
	fmt.Println()
}
