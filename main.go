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
	arr = sort.Insertion(arr)

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
	var s ds.Stack

	fmt.Println(s.Read())

	for i := 0; i < len(arr); i++ {
		s.Push(arr[i])
	}

	fmt.Println(s.Read())
	fmt.Println(s.Pop())
	fmt.Println(s.Read())
}
