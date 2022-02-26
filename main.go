package main

import (
	"fmt"

	"github.com/wolfinger/algos/ds"
	"github.com/wolfinger/algos/search"
	"github.com/wolfinger/algos/sort"
)

func main() {
	arr := []int{24, 3, 7, 6, 8, 103, 18, 1, 52, 12}

	// test bubble sort
	// arr = sort.Bubble(arr)

	// test selection sort
	// arr = sort.Selection(arr)

	// test insertion sort
	// arr = sort.Insertion(arr)

	// test quickselect
	fmt.Println(sort.Quickselect(arr, 2))
	fmt.Println(sort.Quickselect(arr, 0))
	fmt.Println(sort.Quickselect(arr, 4))

	// test quicksort sort
	arr = sort.Quicksort(arr)

	fmt.Println(arr)

	// test binary search
	idx, found := search.Binary(arr, 6)
	fmt.Println(idx, found)

	idx, found = search.Binary(arr, 52)
	fmt.Println(idx, found)

	idx, found = search.Binary(arr, 1)
	fmt.Println(idx, found)

	idx, found = search.Binary(arr, 2)
	fmt.Println(idx, found)

	// test a stack
	fmt.Println("testing stack:")
	var s ds.Stack

	fmt.Println(s.Read())

	for i := 0; i < len(arr); i++ {
		s.Push(arr[i])
	}

	fmt.Println(s.Read())
	fmt.Println(s.Pop())
	fmt.Println(s.Read())

	// test a queue
	fmt.Println("testing queue:")
	var q ds.Queue

	fmt.Println(q.Read())

	for i := 0; i < len(arr); i++ {
		q.Enqueue(arr[i])
	}

	fmt.Println(q.Read())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Read())
}
