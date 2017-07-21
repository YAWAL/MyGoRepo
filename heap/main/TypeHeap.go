package main

import "container/heap"

type Type struct{}
type TypeHeap struct {
	heap typeHeap
}

func (t *TypeHeap) Init(less func(i, j Type) bool) {
	t.heap.less = less
}

func (h *TypeHeap) Push(v Type) {
	heap.Push(&h.heap, v)
}

func (h *TypeHeap) Peek() Type {
	if h.Len() == 0 {
		panic("heap is empty")
	}
	return h.heap.slice[0]
}

func (h *TypeHeap) Pop() Type {
	val, ok := heap.Pop(&h.heap).(Type)
	if !ok {
		panic("invalid type in our heap - this shouldn't ever happen")
	}
	return val
}

func (h *TypeHeap) Len() int {
	return h.heap.Len()
}

type typeHeap struct {
	slice []Type
	less  func(i, j Type) bool
}

func (h typeHeap) Len() int {
	return len(h.slice)
}

func (h typeHeap) Less(i, j int) bool {
	return h.less(h.slice[i], h.slice[j])
}

func (h typeHeap) Swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
}

func (h *typeHeap) Push(x interface{}) {
	h.slice = append(h.slice, x.(Type))
}

func (h *typeHeap) Pop() interface{} {
	n := len(h.slice)
	ret := h.slice[n-1]
	h.slice = h.slice[0 : n-1]
	return ret
}
