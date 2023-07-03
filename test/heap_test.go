package container

import (
	"math/rand"
	"testing"

	"github.com/ralphexp/container"
)

type TestType struct {
	v int
}

func (tt TestType) CompareTo(other interface{}) bool {
	return tt.v < (other.(TestType)).v
}

func TestSequence(t *testing.T) {
	h := container.NewHeap()
	for i := 100; i > 0; i-- {
		h.Push(i)
	}

	a := make([]int, 0)
	for {
		if h.Len() == 0 {
			break
		}
		a = append(a, h.Pop().(int))
	}

	for i := 0; i < 100; i++ {
		if a[i] != i+1 {
			t.Errorf("heap error: expected %d but got %d\n", i+1, a[i])
		}
	}
}

func TestRandomInsert(t *testing.T) {
	a := make([]interface{}, 0)
	for i := 100; i > 0; i-- {
		a = append(a, i)
	}

	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})

	t.Logf("shuffled: %v\n", a)
	h := container.NewHeap()
	h.Init(a)
	a = []interface{}{} // clear a

	for {
		if h.Len() == 0 {
			break
		}
		a = append(a, h.Pop())
	}

	for i := 0; i < 100; i++ {
		if a[i].(int) != i+1 {
			t.Errorf("heap error: expected %d but got %d\n", i+1, a[i])
		}
	}
}

func TestUserDefinedType(t *testing.T) {
	a := make([]int, 0)
	for i := 100; i > 0; i-- {
		a = append(a, i)
	}

	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})

	u := make([]interface{}, 0)
	for i := 0; i < 100; i++ {
		u = append(u, TestType{v: a[i]})
	}

	h := container.NewHeap()
	h.Init(u)
	u = []interface{}{} // clear u

	for {
		if h.Len() == 0 {
			break
		}
		u = append(u, h.Pop())
	}

	for i := 0; i < 100; i++ {
		if (u[i].(TestType)).v != i+1 {
			t.Errorf("heap error: expected %d but got %v\n", i+1, u[i])
		}
	}
}
