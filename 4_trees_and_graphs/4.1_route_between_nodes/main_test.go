package main

import (
	"testing"

	directedgraph "github.com/RolloCasanova/ctci-go/4_trees_and_graphs/utils/directed_graph"
)

func TestRouteBetweenNodes(t *testing.T) {
	type args struct {
		g       directedgraph.DirectedGraph
		start   string
		end     string
		visited map[string]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Direct route exists between nodes",
			args: args{
				g: func() directedgraph.DirectedGraph {
					dg := directedgraph.New()
					dg.AddEdge("A", "B")

					return dg
				}(),
				start:   "A",
				end:     "B",
				visited: make(map[string]bool),
			},
			want: true,
		},
		{
			name: "Short indirect route exists between nodes",
			args: args{
				g: func() directedgraph.DirectedGraph {
					dg := directedgraph.New()
					dg.AddEdge("A", "B")
					dg.AddEdge("B", "C")

					return dg
				}(),
				start:   "A",
				end:     "C",
				visited: make(map[string]bool),
			},
			want: true,
		},
		{
			name: "Long indirect route exists between nodes",
			args: args{
				g: func() directedgraph.DirectedGraph {
					dg := directedgraph.New()
					dg.AddEdge("A", "B")
					dg.AddEdge("B", "C")
					dg.AddEdge("C", "D")
					dg.AddEdge("D", "E")
					dg.AddEdge("E", "F")

					return dg
				}(),
				start:   "A",
				end:     "F",
				visited: make(map[string]bool),
			},
			want: true,
		},
		{
			name: "No route exists between nodes",
			args: args{
				g: func() directedgraph.DirectedGraph {
					dg := directedgraph.New()
					dg.AddEdge("A", "B")
					dg.AddEdge("B", "C")
					dg.AddEdge("C", "D")
					dg.AddEdge("D", "E")
					dg.AddEdge("E", "F")

					return dg
				}(),
				start:   "A",
				end:     "G",
				visited: make(map[string]bool),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RouteBetweenNodes(tt.args.g, tt.args.start, tt.args.end, tt.args.visited); got != tt.want {
				t.Errorf("RouteBetweenNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
