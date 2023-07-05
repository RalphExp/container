package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Value int

func (v Value) Less(other interface{}) bool {
	return v < (other.(Value))
}

func TestInsert0(t *testing.T) {
	tree := NewRBTree()
	size := 100

	for i := 0; i < size; i++ {
		tree.Insert(Value(i), i)
	}

	// tree.Dump()
	assert.Equal(t, tree.len, size)
}

func TestMin0(t *testing.T) {
	tree := NewRBTree()
	for i := 1; i <= 100; i++ {
		tree.Insert(Value(i), i)
	}
	n := tree.Min()
	i := 1
	for n != nil {
		assert.Equal(t, n.Key, Value(i))
		n = tree.Next(n)
		i += 1
	}
}

func TestMax0(t *testing.T) {
	tree := NewRBTree()
	for i := 1; i <= 100; i++ {
		tree.Insert(Value(i), i)
	}
	n := tree.Max()
	i := 100
	for n != nil {
		assert.Equal(t, n.Key, Value(i))
		n = tree.Prev(n)
		i -= 1
	}
}

func TestFind0(t *testing.T) {
	tree := NewRBTree()
	for i := 1; i <= 100; i++ {
		tree.Insert(Value(i), i)
	}

	for i := 1; i <= 100; i++ {
		n := tree.Find(Value(i))
		assert.Equal(t, n.Key, Value(i))
	}
}

func TestFind1(t *testing.T) {
	tree := NewRBTree()
	for i := 1; i <= 100; i++ {
		tree.Insert(Value(i), i)
	}

	// test not found
	n := tree.Find(Value(101))
	assert.Equal(t, n, (*RBNode)(nil))
}

func TestDelete0(t *testing.T) {
	tree := NewRBTree()
	for i := 1; i <= 100; i++ {
		tree.Insert(Value(i), i)
	}

	for i := 1; i <= 100; i++ {
		tree.Delete(Value(i))
	}
	assert.Equal(t, tree.Len(), 0)
}

func TestDelete1(t *testing.T) {
	tree := NewRBTree()
	for i := 1; i <= 100; i++ {
		tree.Insert(Value(i), i)
	}

	for i := 51; i <= 100; i++ {
		tree.Delete(Value(i))
	}
	assert.Equal(t, tree.Len(), 50)

	n := tree.Min()
	for i := 1; i <= 50; i++ {
		assert.Equal(t, n.Key, Value(i))
		assert.Equal(t, n.Value, i)
		n = tree.Next(n)
	}
}
