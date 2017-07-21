package main

import (
	"fmt"
	"container/heap"
)

func main() {
	less := func(i, j string) bool {
		return i < j
	}
	var h String
	h.Init(less)
	h.Push("cat")
	h.Push("dog")
	h.Push("a")
	h.Push("bird with red")
	h.Push("bird")
	for h.Len() > 0 {
		fmt.Println(h.Pop())
	}
}

type String struct {
	heap stringHeap
}

func (t *String) Init(less func(i, j string) bool) {
	t.heap.less = less
}

func (h *String) Push(v string) {
	heap.Push(&h.heap, v)
}

func (h *String) Peek() string {
	if h.Len() == 0 {
		panic("heap is empty")
	}
	return h.heap.slice[0]
}

func (h *String) Pop() string {
	val, ok := heap.Pop(&h.heap).(string)
	if !ok {
		panic("invalid type in our heap - this shouldn't ever happen")
	}
	return val
}

func (h *String) Len() int {
	return h.heap.Len()
}

type stringHeap struct {
	slice []string
	less  func(i, j string) bool
}

func (h stringHeap) Len() int {
	return len(h.slice)
}

func (h stringHeap) Less(i, j int) bool {
	return h.less(h.slice[i], h.slice[j])
}

func (h stringHeap) Swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
}

func (h *stringHeap) Push(x interface{}) {
	h.slice = append(h.slice, x.(string))
}

func (h *stringHeap) Pop() interface{} {
	n := len(h.slice)
	ret := h.slice[n-1]
	h.slice = h.slice[0: n-1]
	return ret
}
