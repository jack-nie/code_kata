package main

import "fmt"

func main() {
	BuildHeap()
}

type Heap struct {
	Capacity int
	Size     int
	Elements []int
}

func InitHeap(maxElements int) *Heap {
	heap := new(Heap)
	heap.Capacity = maxElements + 1
	heap.Size = 0
	heap.Elements = make([]int, maxElements+1)
	heap.Elements[0] = -1
	return heap
}

func (heap *Heap) Insert(x int) {
	var i int
	for i = heap.Size + 1; heap.Elements[i/2] > x; i /= 2 {
		heap.Elements[i] = heap.Elements[i/2]
	}
	heap.Elements[i] = x
	heap.Size++
}

func (heap *Heap) DelelteMin() int {
	lastElement := heap.Elements[heap.Size-1]
	minElement := heap.Elements[1]
	var i, child int
	for i = 1; i*2 <= heap.Size; i = child {
		child = i * 2
		if child != heap.Size && heap.Elements[child] > heap.Elements[child+1] {
			child++
		}
		if lastElement > heap.Elements[child] {
			heap.Elements[i] = heap.Elements[child]
		} else {
			break
		}
	}
	heap.Elements[i] = lastElement
	heap.Elements = heap.Elements[:heap.Size]
	fmt.Println(heap.Elements[:])
	heap.Size--
	return minElement
}

func BuildHeap() {
	heap := InitHeap(5)
	for _, item := range []int{4, 2, 8, 9, 1} {
		heap.Insert(item)
	}
	fmt.Println(heap.Elements[1:])
	i := 5
	for i > 0 {
		fmt.Println(heap.DelelteMin())
		i--
	}
}
