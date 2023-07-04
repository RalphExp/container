// heap.go use the builtin container/heap's algorithm
// but make it more object-orient and easier to use,
// user only need to implement the comparable interface to
// utilize the Heap struct

package container

type Heap struct {
	heap []Comparable
}

func NewHeap() *Heap {
	return &Heap{heap: make([]Comparable, 0)}
}

func (h *Heap) Len() int {
	return len(h.heap)
}

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
func (h *Heap) Init(values []Comparable) {
	// heapify
	h.heap = values
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Push(x Comparable) {
	h.heap = append(h.heap, x)
	h.up(h.Len() - 1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func (h *Heap) Pop() Comparable {
	if h.Len() == 0 {
		return nil
	}
	n := h.Len() - 1
	h.Swap(0, n)
	h.down(0, n)
	ret := h.heap[n]
	h.heap = h.heap[0:n]
	return ret
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Fix(i int) {
	if !h.down(i, h.Len()) {
		h.up(i)
	}
}

func (h Heap) Swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap) Remove(i int) any {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !h.down(i, n) {
			h.up(i)
		}
	}
	return h.Pop()
}

// Get the underlying slice
func (h *Heap) GetSlice() []Comparable {
	return h.heap
}

func (h Heap) less(i, j int) bool {
	return h.heap[i].Less(h.heap[j])
}

func (h *Heap) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func (h *Heap) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
