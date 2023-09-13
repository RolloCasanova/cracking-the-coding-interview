package directedgraph

import "slices"

// DirectedGraph is the interface for a directed graph
type DirectedGraph interface {
	AddNode(value string)
	AddEdge(from, to string)
	GetNode(value string) *node
	RemoveNode(value string)
	RemoveEdge(from, to string)
	HasNode(value string) bool
	HasEdge(from, to string) bool
}

// directedGraph is the implementation of DirectedGraph
type directedGraph struct {
	Nodes []*node
}

// node is a node in the graph
type node struct {
	Value string
	Edges []*node
}

// New returns a new instance of *directedGraph
func New() DirectedGraph {
	return &directedGraph{}
}

// AddNode adds a node to the graph
func (g *directedGraph) AddNode(value string) {
	if g.HasNode(value) {
		return
	}

	g.Nodes = append(g.Nodes, &node{Value: value, Edges: make([]*node, 0)})
}

// AddEdge adds an edge to the graph from "from" node to "to" node
// if the nodes do not exist, they are created
func (g *directedGraph) AddEdge(from, to string) {
	fromNode := g.GetNode(from)
	if fromNode == nil {
		fromNode = &node{
			Value: from,
		}
	}

	toNode := g.GetNode(to)
	if toNode == nil {
		toNode = &node{
			Value: to,
		}
	}

	if !slices.Contains(g.Nodes, fromNode) {
		g.Nodes = append(g.Nodes, fromNode)
	}

	if !slices.Contains(g.Nodes, toNode) {
		g.Nodes = append(g.Nodes, toNode)
	}

	if !slices.Contains(fromNode.Edges, toNode) {
		fromNode.Edges = append(fromNode.Edges, toNode)
	}
}

// RemoveNode removes a node from the graph and all edges to it
// if node does not exist, nothing happens
func (g *directedGraph) RemoveNode(value string) {
	node := g.GetNode(value)
	if node == nil {
		return
	}

	// remove the node from the graph
	if index := slices.Index(g.Nodes, node); index != -1 {
		slices.Delete(g.Nodes, index, index+1)
	}

	// remove all edges to the node
	for _, n := range g.Nodes {
		if index := slices.Index(n.Edges, node); index != -1 {
			slices.Delete(n.Edges, index, index+1)
		}
	}
}

// RemoveEdge removes an edge from the graph from "from" node to "to" node
// if a node or edge do not exist, nothing happens
func (g *directedGraph) RemoveEdge(from, to string) {
	fromNode := g.GetNode(from)
	if fromNode == nil {
		return
	}

	toNode := g.GetNode(to)
	if toNode == nil {
		return
	}

	if slices.Contains(fromNode.Edges, toNode) {
		fromNode.Edges = append(fromNode.Edges[:0], fromNode.Edges[1:]...)
		return
	}
}

// HasNode returns true if the graph has a node with the given value
func (g *directedGraph) HasNode(value string) bool {
	return g.GetNode(value) != nil
}

// HasEdge returns true if there is an edge from "from" node to "to" node
// if a node or edge do not exist, returns false
func (g *directedGraph) HasEdge(from, to string) bool {
	fromNode := g.GetNode(from)
	if fromNode == nil {
		return false
	}

	toNode := g.GetNode(to)
	if toNode == nil {
		return false
	}

	return slices.Contains(fromNode.Edges, toNode)
}

// GetNode returns the node with the given value
func (g *directedGraph) GetNode(value string) *node {
	if len(g.Nodes) == 0 {
		return nil
	}

	for _, n := range g.Nodes {
		if n.Value == value {
			return n
		}
	}

	return nil
}
