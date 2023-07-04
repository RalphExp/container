// rbtree.go implements an red-black search tree which
// the standard library doest not have

package container

// Red-Black Tree
type RBTree struct {
	root *RBNode
	null *RBNode // sentinel
	len  int
}

// Tree Node
type RBNode struct {
	left   *RBNode
	right  *RBNode
	parent *RBNode
	color  int
	Key    Comparable
	Value  interface{}
}

const (
	red = iota
	black
)

func NewRBTree() *RBTree {
	tree := &RBTree{null: &RBNode{color: black}}
	// set sentienal's parent, left and right to null,
	// they can be set to any values, but null can
	// make it more easier to detect error
	tree.null.parent = nil
	tree.null.left = nil
	tree.null.right = nil
	tree.root = tree.null
	return tree
}

// Insert a new node into the tree, if the key already exists,
// the value will be replace by the new one.
func (tree *RBTree) Insert(key Comparable, value interface{}) {
	z := &RBNode{Key: key, Value: value}
	tree.insertNode(z)
}

// ITA: chap13.3
func (tree *RBTree) insertNode(z *RBNode) {
	var x *RBNode = tree.root
	var y *RBNode = tree.null
	for x != tree.null {
		y = x
		if z.Key.Less(x.Key) {
			x = x.left
		} else {
			x = x.right
		}
	}
	// now x must be tree.null
	if y == tree.null {
		tree.root = z
	} else { // y != nil
		if z.Key.Less(y.Key) {
			y.left = z
		} else if y.Key.Less(z.Key) {
			y.right = z
		} else {
			// replace, the key can be different with y.key
			// depend on how Less works.
			y.Key = z.Key
			y.Value = z.Value
			return
		}
	}
	z.parent = y // y may be tree.null
	z.color = red
	z.left = tree.null
	z.right = tree.null
	tree.insertFixup(z)
	tree.len += 1
}

// ITA: chap13.3
func (tree *RBTree) insertFixup(z *RBNode) {
	for z.parent.color == red {
		// sentinel is used, so the following statement is valid
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			// case 1:
			if y.color == red {
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			} else {
				// case 2: y.color == black && z is the right child
				if z == z.parent.right {
					z = z.parent
					tree.rotateLeft(z)
				}
				// case 3: y.color == black && z is the left child
				z.parent.color = black
				z.parent.parent.color = red
				tree.rotateRight(z.parent.parent)
			}
		} else {
			// z.parent == z.parent.parent.right must be true
			// because the color of z.parent is red, z.parent must
			// have a parent.
			y := z.parent.parent.left
			if y.color == red {
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					tree.rotateRight(z)
				}
				z.parent.color = black
				z.parent.parent.color = red
				tree.rotateLeft(z.parent.parent)
			}
		}
	}
	tree.root.color = black
}

func (tree *RBTree) Remove(key Comparable) bool {
	return false
}

// Find return a *copy* of the Red-Black Tree node
// or nil if the given key does not exist in the tree.
// Note that although you can change the key of the
// node, the corresponding node of the tree is intact
// because the return value is just a copy.
func (tree *RBTree) Find(key Comparable) *RBNode {
	return nil
}

func (tree *RBTree) Len() int {
	return tree.len
}

func (tree *RBTree) Begin() *RBNode {
	return nil
}

func (node *RBNode) Next() *RBNode {
	return nil
}

func (node *RBNode) Prev() *RBNode {
	return nil
}

// ITA chap13.2
func (tree *RBTree) rotateLeft(x *RBNode) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		tree.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

// Dual of rotateLeft (change left -> right and right -> left)
func (tree *RBTree) rotateRight(x *RBNode) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		tree.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}
