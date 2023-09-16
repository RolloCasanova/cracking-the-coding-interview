package bst

// BST is a binary search tree
type BST struct {
	Value int
	Left  *BST
	Right *BST
}

// New returns a new instance of *BST with the given values inserted
func New(values ...int) *BST {
	var bst *BST

	for _, value := range values {
		bst = bst.Insert(value)
	}

	return bst
}

// Insert adds a value to the BST and returns the root node
func (t *BST) Insert(value int) *BST {
	if t == nil {
		return &BST{Value: value}
	}

	if value <= t.Value {
		t.Left = t.Left.Insert(value)
	} else {
		t.Right = t.Right.Insert(value)
	}

	return t
}
