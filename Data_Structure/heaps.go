package main

import (
	"container/heap"
	"fmt"
)

// An IntegerHeap is a min-heap of ints.
type IntegerHeap []int

// IntegerHeap implements heap.Interface and holds an int slice.
func (h IntegerHeap) Len() int {
	return len(h)
}

// Less returns whether the element with index i should sort before the element with index j.
func (h IntegerHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap swaps the elements with indexes i and j.
func (h IntegerHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push pushes the element x onto the heap.
func (h *IntegerHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes the minimum element (according to Less) from the heap and returns it.
func (h *IntegerHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// main function
func main() {
	h := &IntegerHeap{1, 4, 5}
	heap.Init(h)
	heap.Push(h, 2)
	fmt.Printf("minimum: %d", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
