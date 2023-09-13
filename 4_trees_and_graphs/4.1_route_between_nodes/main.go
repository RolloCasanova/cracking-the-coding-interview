package main

import (
	"fmt"
	"os"

	directedgraph "github.com/RolloCasanova/ctci-go/4_trees_and_graphs/utils/directed_graph"

	"github.com/emirpasic/gods/queues/arrayqueue"
)

func main() {
	// <nodeA> <nodeB> represents two nodes values in the graph, and the program will determine if there is a route between them
	// <start_node> <end_node> creates a directed edge from <start_node> to <end_node> to the graph
	// <start_node> 0 will add a node with value <start_node> and no edges to the graph
	if len(os.Args) < 3 || len(os.Args[3:])%2 != 0 {
		panic("usage: go run main.go <nodeA> <nodeB> [<start_node> [<end_node>|0]]...")
	}

	g := directedgraph.New()

	for i := 3; i < len(os.Args); i += 2 {
		if os.Args[i+1] == "0" {
			g.AddNode(os.Args[i])
		} else {
			g.AddEdge(os.Args[i], os.Args[i+1])
		}
	}

	var operation, nodeA, nodeB string
loop:
	for {
		fmt.Println("enter operation to perform on graph (add_node, add_edge, remove_node, remove_edge, has_node, has_edge), or exit to quit:")
		_, err := fmt.Scan(&operation)
		if err != nil {
			break
		}

		// if the operation is enqueue, we need to read the animal kind and name
		switch operation {
		case "add_node":
			fmt.Println("enter node value: ")
			_, err := fmt.Scan(&nodeA)
			if err != nil {
				fmt.Println("node value is required")
				break
			}

			g.AddNode(nodeA)
		case "add_edge":
			fmt.Println("enter nodeA value: ")
			_, err := fmt.Scan(&nodeA)
			if err != nil {
				fmt.Println("nodeA value is required")
				break
			}

			fmt.Println("enter nodeB value: ")
			_, err = fmt.Scan(&nodeB)
			if err != nil {
				fmt.Println("nodeB value is required")
				break
			}

			g.AddEdge(nodeA, nodeB)

		case "remove_node":
			fmt.Println("enter node value: ")
			_, err := fmt.Scan(&nodeA)
			if err != nil {
				fmt.Println("node value is required")
				break
			}

			g.RemoveNode(nodeA)

		case "remove_edge":
			fmt.Println("enter nodeA value: ")
			_, err := fmt.Scan(&nodeA)
			if err != nil {
				fmt.Println("nodeA value is required")
				break
			}

			fmt.Println("enter nodeB value: ")
			_, err = fmt.Scan(&nodeB)
			if err != nil {
				fmt.Println("nodeB value is required")
				break
			}

			g.RemoveEdge(nodeA, nodeB)

		case "has_node":
			fmt.Println("enter node value: ")
			_, err := fmt.Scan(&nodeA)
			if err != nil {
				fmt.Println("node value is required")
				break
			}

			if g.HasNode(nodeA) {
				fmt.Println("Node " + nodeA + " exists")
			} else {
				fmt.Println("Node " + nodeA + " does not exist")
			}

		case "has_edge":
			fmt.Println("enter nodeA value: ")
			_, err := fmt.Scan(&nodeA)
			if err != nil {
				fmt.Println("nodeA value is required")
				break
			}

			fmt.Println("enter nodeB value: ")
			_, err = fmt.Scan(&nodeB)
			if err != nil {
				fmt.Println("nodeB value is required")
				break
			}

			if g.HasEdge(nodeA, nodeB) {
				fmt.Println("Edge from " + nodeA + " to " + nodeB + " exists")
			} else {
				fmt.Println("Edge from " + nodeA + " to " + nodeB + " does not exist")
			}

		case "exit":
			fmt.Println()
			break loop

		default:
			fmt.Println("invalid operation \" " + operation + "\"")
		}

		fmt.Println()
	}

	if RouteBetweenNodes(g, os.Args[1], os.Args[2], make(map[string]bool)) {
		println("There is a route between", os.Args[1], "and", os.Args[2])
	} else {
		println("There is no route between", os.Args[1], "and", os.Args[2])
	}
}

func RouteBetweenNodes(g directedgraph.DirectedGraph, start, end string, visited map[string]bool) bool {
	q := arrayqueue.New()
	visited[start] = true
	q.Enqueue(start)

	for !q.Empty() {
		v, _ := q.Dequeue()
		s := v.(string)

		if g.HasEdge(s, end) {
			return true
		}

		n := g.GetNode(s)

		for _, node := range n.Edges {
			if !visited[node.Value] {
				q.Enqueue(node.Value)
			}
		}
	}

	return false
}
