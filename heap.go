package main

// KeySize contains a key and its size.
type KeySize struct {
	key  string
	size int64
}

// Heap implements container/heap interface.
type Heap []*KeySize

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) Less(i, j int) bool {
	if h[i].size == h[j].size {
		return h[i].key < h[j].key
	}
	return h[i].size > h[j].size
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*KeySize))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
