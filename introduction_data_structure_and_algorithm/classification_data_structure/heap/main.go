package main

import (
	"container/heap"
	"fmt"
)

type IntegerHeap []int

func (a IntegerHeap) Len() int           { return len(a) }
func (a IntegerHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntegerHeap) Less(i, j int) bool { return a[i] < a[j] }

func (a *IntegerHeap) Push(heapintf interface{}) {
	*a = append(*a, heapintf.(int))
}

func (a *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int

	previous := *a

	n = len(previous)
	x1 = previous[n-1]
	*a = previous[0 : n-1]

	return x1
}

func main() {
	intHeap := &IntegerHeap{1, 2, 3}

	heap.Init(intHeap)
	heap.Push(intHeap, 4)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])

	for intHeap.Len() > 0 {
		fmt.Printf("%d\n", heap.Pop(intHeap))
	}

	fmt.Printf("intheap len after pop iterration: %d\n", intHeap.Len())
}
