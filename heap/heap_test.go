package heap

import (
	"math/rand"
	"testing"

	"github.com/ralphexp/container"
)

type Value int

func (v Value) Less(other interface{}) bool {
	return v < (other.(Value))
}

func TestRandomInsert(t *testing.T) {
	a := make([]container.Comparable, 0)
	for i := 0; i < 1000; i++ {
		a = append(a, Value(i))
	}

	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})

	h := NewHeap()
	h.Init(a)
	a = []container.Comparable{} // clear a

	for {
		if h.Len() == 0 {
			break
		}
		a = append(a, h.Pop())
	}

	for i := 0; i < 1000; i++ {
		if a[i].(Value) != Value(i) {
			t.Errorf("heap error: expected %d but got %d\n", i, a[i])
		}
	}
}

func TestFix(t *testing.T) {
	a := make([]container.Comparable, 0)
	for i := 0; i < 100; i++ {
		a = append(a, Value(i))
	}

	h := NewHeap()
	h.Init(a)
	for i := 0; i < 10; i++ {
		h.GetSlice()[i*10] = Value(i * 10)
		h.Fix(i * 10)
	}

	j := container.Comparable(Value(-1))
	for i := 0; i < 100; i++ {
		k := h.Pop()
		if k.Less(j) {
			t.Errorf("heap error: j = %d, k = %d\n", j, k)
		}
		j = k
	}
}
