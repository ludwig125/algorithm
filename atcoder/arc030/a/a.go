package main

import "fmt"

func main() {
	solve()
}

func solve() {
	var n, k int
	fmt.Scan(&n, &k)

	g := NewGraph(n)
	fmt.Printf("g %v\n", g)
}

// Graph is struct.
type Graph struct {
	Nodes []Node
}

// Node is struct.
type Node struct {
	ID    int
	Nodes []Node
}

// NewNode makes new Node.
func NewNode(id int, ids ...int) *Node {
	n := Node{ID: id}
	for _, i := range ids {
		n.addNode(i)
	}
	return &n
}

func (n *Node) addNode(id int) {
	newNode := Node{ID: id}
	n.Nodes = append(n.Nodes, newNode)
}

func (n *Node) nodesNum(id int) int {
	count := 1

	return count
}

func traverse(n Node, seenNodes []int) {
	for _,v :=range n.Nodes{
		if v.ID
	}
}

// NewGraph makes new Graph.
func NewGraph(n int) *Graph {
	var nodes []Node
	for i := 1; i <= n; i++ {
		nodes = append(nodes, *NewNode(i, nextID(n, i), prevID(n, i)))
	}
	return &Graph{
		Nodes: nodes,
	}
}

func nextID(n, id int) int {
	if id == n {
		return 1
	}
	return id + 1
}

func prevID(n, id int) int {
	if id == 1 {
		return n
	}
	return id - 1
}
