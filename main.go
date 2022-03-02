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

	// test a linked list
	n1 := ds.NewNode(arr[0])
	n2 := ds.NewNode(arr[1])
	n3 := ds.NewNode(arr[3])

	ll := ds.NewLinkedList(n1)
	ll.AddToEnd(n2)
	ll.AddToEnd(n3)

	// print nodes
	for node := ll.Head; node != nil; node = node.Next {
		fmt.Println(node.Data)
	}

	ll.DeleteFromEnd()

	// print nodes using read method
	fmt.Println(ll.Read(0))
	fmt.Println(ll.Read(1))
	fmt.Println(ll.Read(2))

	/*
		// test a bst
		bstnR1N1 := ds.NewBSTNode(12) //                  12
		bstnR2N1 := ds.NewBSTNode(7)  //               7      24
		bstnR2N2 := ds.NewBSTNode(24) //             3   8  18   52
		bstnR3N1 := ds.NewBSTNode(3)  //           1  6           103
		bstnR3N2 := ds.NewBSTNode(8)
		bstnR3N3 := ds.NewBSTNode(18)
		bstnR3N4 := ds.NewBSTNode(52)
		bstnR4N1 := ds.NewBSTNode(1)
		bstnR4N2 := ds.NewBSTNode(6)
		bstnR4N3 := ds.NewBSTNode(103)

		// first level
		bstnR1N1.Left = bstnR2N1
		bstnR1N1.Right = bstnR2N2

		// second level
		bstnR2N1.Left = bstnR3N1
		bstnR2N1.Right = bstnR3N2
		bstnR2N2.Left = bstnR3N3
		bstnR2N2.Right = bstnR3N4

		// third level
		bstnR3N1.Left = bstnR4N1
		bstnR3N1.Right = bstnR4N2
		bstnR3N4.Right = bstnR4N3

		// set root alias
		root := bstnR1N1
	*/

	root := ds.NewBSTNode(12)
	ds.BSTInsert(root, 7)
	ds.BSTInsert(root, 24)
	ds.BSTInsert(root, 3)
	ds.BSTInsert(root, 8)
	ds.BSTInsert(root, 18)
	ds.BSTInsert(root, 52)
	ds.BSTInsert(root, 1)
	ds.BSTInsert(root, 6)
	ds.BSTInsert(root, 103)

	fmt.Println(ds.BSTSearch(root, 18))
	fmt.Println(ds.BSTSearch(root, 3))
	fmt.Println(ds.BSTSearch(root, 21))
}
