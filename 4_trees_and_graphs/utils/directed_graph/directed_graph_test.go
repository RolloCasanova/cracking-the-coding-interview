package directedgraph

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want DirectedGraph
	}{
		{
			name: "Create a new directed graph",
			want: &directedGraph{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_directedGraph_AddNode(t *testing.T) {
	type fields struct {
		Nodes []*node
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Add a non-existing node to the graph",
			fields: fields{
				Nodes: []*node{
					{Value: "A"},
				},
			},
			args: args{
				value: "B",
			},
		},
		{
			name: "Add an existing node to the graph",
			fields: fields{
				Nodes: []*node{
					{Value: "A"},
				},
			},
			args: args{
				value: "A",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			g := &directedGraph{
				Nodes: tt.fields.Nodes,
			}
			g.AddNode(tt.args.value)
		})
	}
}

func Test_directedGraph_AddEdge(t *testing.T) {
	type fields struct {
		Nodes []*node
	}
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Add an edge between two existing nodes",
			fields: fields{
				Nodes: []*node{
					{Value: "A"},
					{Value: "B"},
				},
			},
			args: args{
				from: "A",
				to:   "B",
			},
		},
		{
			name: "Add an edge between a non-existing node and an existing node",
			fields: fields{
				Nodes: []*node{
					{Value: "A"},
				},
			},
			args: args{
				from: "B",
				to:   "A",
			},
		},
		{
			name: "Add an edge between an existing node and a non-existing node",
			fields: fields{
				Nodes: []*node{
					{Value: "A"},
				},
			},
			args: args{
				from: "A",
				to:   "B",
			},
		},
		{
			name: "Add an edge between two non-existing nodes",
			fields: fields{
				Nodes: []*node{},
			},
			args: args{
				from: "A",
				to:   "B",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			g := &directedGraph{
				Nodes: tt.fields.Nodes,
			}
			g.AddEdge(tt.args.from, tt.args.to)
		})
	}
}

func Test_directedGraph_RemoveNode(t *testing.T) {
	type fields struct {
		Nodes []*node
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Remove an existing node from the graph",
			fields: fields{
				Nodes: []*node{
					{Value: "A"},
				},
			},
			args: args{
				value: "A",
			},
		},
		{
			name: "Remove a non-existing node from the graph",
			fields: fields{
				Nodes: []*node{
					{Value: "A"},
				},
			},
			args: args{
				value: "B",
			},
		},
		{
			name: "Remove edges to the node",
			fields: fields{
				Nodes: func() []*node {
					a := auxCreateNode("A")
					b := auxCreateNode("B")
					a.Edges = []*node{b}

					return []*node{a, b}
				}(),
			},
			args: args{
				value: "B",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			g := &directedGraph{
				Nodes: tt.fields.Nodes,
			}
			g.RemoveNode(tt.args.value)
		})
	}
}

func Test_directedGraph_RemoveEdge(t *testing.T) {
	type fields struct {
		Nodes []*node
	}
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Remove an existing edge from the graph",
			fields: fields{
				Nodes: func() []*node {
					a := auxCreateNode("A")
					b := auxCreateNode("B")
					a.Edges = []*node{b}

					return []*node{a, b}
				}(),
			},
			args: args{
				from: "A",
				to:   "B",
			},
		},
		{
			name: "Remove a non-existing edge from the graph",
			fields: fields{
				Nodes: func() []*node {
					a := auxCreateNode("A")
					b := auxCreateNode("B")
					a.Edges = []*node{b}

					return []*node{a, b}
				}(),
			},
			args: args{
				from: "B",
				to:   "A",
			},
		},
		{
			name: "Remove an edge from a non-existing node",
			fields: fields{
				Nodes: func() []*node {
					a := auxCreateNode("A")
					b := auxCreateNode("B")
					a.Edges = []*node{b}

					return []*node{a, b}
				}(),
			},
			args: args{
				from: "C",
				to:   "A",
			},
		},
		{
			name: "Remove an edge to a non-existing node",
			fields: fields{
				Nodes: func() []*node {
					a := auxCreateNode("A")
					b := auxCreateNode("B")
					a.Edges = []*node{b}

					return []*node{a, b}
				}(),
			},
			args: args{
				from: "A",
				to:   "C",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			g := &directedGraph{
				Nodes: tt.fields.Nodes,
			}
			g.RemoveEdge(tt.args.from, tt.args.to)
		})
	}
}

func Test_directedGraph_HasNode(t *testing.T) {
	type fields struct {
		Nodes []*node
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Node exists",
			fields: fields{
				Nodes: []*node{
					{Value: "A"},
				},
			},
			args: args{
				value: "A",
			},
			want: true,
		},
		{
			name: "Node does not exist",
			fields: fields{
				Nodes: []*node{
					{Value: "A"},
				},
			},
			want: false,
		},
		{
			name: "Empty graph",
			fields: fields{
				Nodes: []*node{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &directedGraph{
				Nodes: tt.fields.Nodes,
			}
			if got := g.HasNode(tt.args.value); got != tt.want {
				t.Errorf("directedGraph.HasNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_directedGraph_HasEdge(t *testing.T) {
	type fields struct {
		Nodes []*node
	}
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Edge exists",
			fields: fields{
				Nodes: func() []*node {
					a := auxCreateNode("A")
					b := auxCreateNode("B")
					a.Edges = []*node{b}

					return []*node{a, b}
				}(),
			},
			args: args{
				from: "A",
				to:   "B",
			},
			want: true,
		},
		{
			name: "Edge does not exist",
			fields: fields{
				Nodes: func() []*node {
					a := auxCreateNode("A")
					b := auxCreateNode("B")
					a.Edges = []*node{b}

					return []*node{a, b}
				}(),
			},
			args: args{
				from: "B",
				to:   "A",
			},
			want: false,
		},
		{
			name: "Edge from non-existing node",
			fields: fields{
				Nodes: func() []*node {
					a := auxCreateNode("A")
					b := auxCreateNode("B")
					a.Edges = []*node{b}

					return []*node{a, b}
				}(),
			},
			args: args{
				from: "C",
				to:   "A",
			},
			want: false,
		},
		{
			name: "Edge to non-existing node",
			fields: fields{
				Nodes: func() []*node {
					a := auxCreateNode("A")
					b := auxCreateNode("B")
					a.Edges = []*node{b}

					return []*node{a, b}
				}(),
			},
			args: args{
				from: "A",
				to:   "C",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &directedGraph{
				Nodes: tt.fields.Nodes,
			}
			if got := g.HasEdge(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("directedGraph.HasEdge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_directedGraph_GetNode(t *testing.T) {
	type fields struct {
		Nodes []*node
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node
	}{
		{
			name: "Node exists",
			fields: fields{
				Nodes: []*node{
					{Value: "A"},
				},
			},
			args: args{
				value: "A",
			},
			want: &node{
				Value: "A",
			},
		},
		{
			name: "Node does not exist",
			fields: fields{
				Nodes: []*node{
					{Value: "A"},
				},
			},
			args: args{
				value: "B",
			},
			want: nil,
		},
		{
			name: "Empty graph",
			fields: fields{
				Nodes: []*node{},
			},
			args: args{
				value: "A",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &directedGraph{
				Nodes: tt.fields.Nodes,
			}
			if got := g.GetNode(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("directedGraph.GetNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func auxCreateNode(value string) *node {
	return &node{
		Value: value,
	}
}
