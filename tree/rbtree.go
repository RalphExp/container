// rbtree.go implements an red-black search tree which
// the standard library doest not have

package tree

import (
	"fmt"

	c "github.com/ralphexp/container"
)

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
	Key    c.Comparable
	Value  interface{}
}

const (
	red = iota
	black
)

func NewRBTree() *RBTree {
	tree := &RBTree{null: &RBNode{color: black}}
	// Sentinel's parent, left and right can be set to any
	// value. (See remove Node). So we don't assign any
	// value at the beginning.
	tree.root = tree.null
	return tree
}

// Insert a new node into the tree, if the key already exists,
// the value will be replace by the new one.
func (tree *RBTree) Insert(key c.Comparable, value interface{}) {
	z := &RBNode{Key: key, Value: value}
	tree.insertNode(z)
}

// ITA: chap13.3
func (tree *RBTree) insertNode(z *RBNode) {
	var y *RBNode = tree.null
	var x *RBNode = tree.root
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
	z.parent = y
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

func (tree *RBTree) Delete(key c.Comparable) bool {
	var n *RBNode = tree.Find(key)
	if n == nil {
		return false
	}
	tree.deleteNode(n)
	tree.len -= 1
	return true
}

// ITA: ch13.4
func (tree *RBTree) deleteNode(z *RBNode) {
	var x, y *RBNode
	if z.left == tree.null || z.right == tree.null {
		y = z
	} else {
		y = tree.Next(z)
	}

	if y.left != tree.null {
		x = y.left
	} else {
		x = y.right
	}

	x.parent = y.parent
	if y.parent == tree.null {
		tree.root = x
	} else {
		if y == y.parent.left {
			y.parent.left = x
		} else {
			y.parent.right = x
		}
	}
	if y != z {
		z.Key = y.Key
		z.Value = y.Value
	}
	if y.color == black {
		// x can be:
		// 1) Tree.nil
		// 2) Left child of y if y's right child is Tree.nil
		// 3) Right child of y if y's left child is Tree.nil
		tree.deleteFixup(x)
	}
}

// ITA ch13.4
func (tree *RBTree) deleteFixup(x *RBNode) {
	// x can be tree.nil!!!
	for x != tree.root && x.color == black {
		if x == x.parent.left {
			w := x.parent.right
			// case 1:
			if w.color == red {
				w.color = black
				x.parent.color = red
				tree.rotateLeft(x.parent)
				w = x.parent.right
			}
			// case 2: w.color == black
			if w.left.color == black && w.right.color == black {
				w.color = red
				x = x.parent
			} else {
				// case 3: w.right.color == black, w.left.color == red
				if w.right.color == black {
					w.left.color = black
					w.color = red
					tree.rotateRight(w)
					w = x.parent.right
				}
				// case 4: w.right.color == red
				w.color = x.parent.color
				x.parent.color = black
				w.right.color = black
				tree.rotateLeft(x.parent)
				// set x = tree.root to break the loop
				x = tree.root
			}
		} else {
			// x == x.parent.right
			w := x.parent.left
			if w.color == red {
				w.color = black
				x.parent.color = red
				tree.rotateRight(x.parent)
				w = x.parent.left
			}
			if w.right.color == black && w.left.color == black {
				w.color = red
				x = x.parent
			} else {
				if w.left.color == black {
					w.right.color = black
					w.color = red
					tree.rotateLeft(w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = black
				w.left.color = black
				tree.rotateRight(x.parent)
				x = tree.root
			}
		}
	}
	x.color = black
}

// Find return a Red-Black Tree node pointer or nil
// if the given key does not exist in the tree.
// XXX: don't change the key! or the tree would be
// probably corrupted.
func (tree *RBTree) Find(key c.Comparable) *RBNode {
	var node *RBNode = tree.root
	for node != tree.null {
		if key.Less(node.Key) {
			node = node.left
		} else if node.Key.Less(key) {
			node = node.right
		} else {
			return node
		}
	}
	if node == tree.null {
		return nil
	}
	return node
}

func (tree *RBTree) Len() int {
	return tree.len
}

func (tree *RBTree) Min() *RBNode {
	var x *RBNode = tree.root
	if x == tree.null {
		return nil
	}

	for x.left != tree.null {
		x = x.left
	}
	return x
}

func (tree *RBTree) Max() *RBNode {
	var x *RBNode = tree.root
	if x == tree.null {
		return nil
	}

	for x.right != tree.null {
		x = x.right
	}
	return x
}

// Returns the succesor
func (tree *RBTree) Next(x *RBNode) *RBNode {
	var y *RBNode
	if x.right != tree.null {
		y = x.right
		for y.left != tree.null {
			y = y.left
		}
	} else { // x.right == nil
		y = x.parent
		for y != tree.null && x == y.right {
			x = y
			y = y.parent
		}
	}
	if y == tree.null {
		return nil
	}
	return y
}

// Returns the predecessor
func (tree *RBTree) Prev(x *RBNode) *RBNode {
	var y *RBNode
	if x.left != tree.null {
		y = x.left
		for y.right != tree.null {
			y = y.right
		}
	} else { // x.left == nil
		y = x.parent
		for y != tree.null && x == y.left {
			x = y
			y = y.parent
		}
	}
	if y == tree.null {
		return nil
	}
	return y
}

// ITA chap13.2
func (tree *RBTree) rotateLeft(x *RBNode) {
	y := x.right
	x.right = y.left
	if y.left != tree.null {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == tree.null {
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
	if y.right != tree.null {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == tree.null {
		tree.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

// for debugging
func (tree *RBTree) Dump() {
	if tree.root == tree.null {
		fmt.Println("null tree")
		return
	}
	tree.dumpNode(tree.root)
	fmt.Printf("\n")
}

func (tree *RBTree) dumpNode(n *RBNode) {
	if n == tree.null {
		return
	}
	fmt.Printf("key: %v, color: %d", n.Key, n.color)
	if n.left != tree.null {
		fmt.Printf(", left: %v", n.left.Key)
	} else {
		fmt.Printf(", left: nil")
	}
	if n.right != tree.null {
		fmt.Printf(", right: %v", n.right.Key)
	} else {
		fmt.Printf(", right: nil")
	}
	fmt.Printf("\n")
	tree.dumpNode(n.left)
	tree.dumpNode(n.right)
}
