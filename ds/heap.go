package ds

import "fmt"

type Heap struct {
	data []int
}

func NewHeap(val int) *Heap {
	return &Heap{
		data: []int{val},
	}
}

func (heap *Heap) Print() {
	for i := 0; i < len(heap.data); i++ {
		fmt.Println(heap.data[i])
	}
}

func (heap *Heap) RootNode() *int {
	if len(heap.data) > 0 {
		return &heap.data[0]
	} else {
		return nil
	}
}

func (heap *Heap) LastNode() *int {
	if len(heap.data) > 0 {
		return &heap.data[len(heap.data)-1]
	} else {
		return nil
	}
}

func (heap *Heap) LeftKid(index int) int {
	return (index * 2) + 1
}

func (heap *Heap) RightKid(index int) int {
	return (index * 2) + 2
}

func (heap *Heap) Parent(index int) int {
	return (index - 1) / 2
}

func (heap *Heap) Insert(val int) *int {
	heap.data = append(heap.data, val)

	valIdx := len(heap.data) - 1
	parentIdx := heap.Parent(valIdx)
	for heap.data[valIdx] > heap.data[parentIdx] {
		heap.data[parentIdx], heap.data[valIdx] = heap.data[valIdx], heap.data[parentIdx]
		valIdx = parentIdx
		parentIdx = heap.Parent(valIdx)
	}

	return &heap.data[valIdx]
}

func (heap *Heap) Delete() int {
	retVal := *heap.RootNode()

	// swap last item with root
	*heap.RootNode() = *heap.LastNode()

	// delete last node
	heap.data = heap.data[:len(heap.data)-1]

	// trickle down the root note
	trickleIdx := 0
	for heap.HasGreaterKid(trickleIdx) {
		largerKidIdx := heap.CalcLargerKid(trickleIdx)
		heap.data[largerKidIdx], heap.data[trickleIdx] = heap.data[trickleIdx], heap.data[largerKidIdx]
		trickleIdx = largerKidIdx
	}

	return retVal
}

func (heap *Heap) HasGreaterKid(index int) bool {
	if (heap.LeftKid(index) < len(heap.data) && (heap.data[index] < heap.data[heap.LeftKid(index)])) ||
		(heap.RightKid(index) < len(heap.data) && (heap.data[index] < heap.data[heap.RightKid(index)])) {
		return true
	} else {
		return false
	}
}

func (heap *Heap) CalcLargerKid(index int) int {
	var leftKid, rightKid int

	if heap.LeftKid(index) < len(heap.data) {
		leftKid = heap.LeftKid(index)
	}

	if heap.RightKid(index) >= len(heap.data) {
		return leftKid
	} else {
		rightKid = heap.RightKid(index)
	}

	if heap.data[rightKid] > heap.data[leftKid] {
		return rightKid
	} else {
		return leftKid
	}
}
