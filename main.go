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

	// test a binary search tree
	//                   12
	//               7       24
	//             3   8   18   52
	//           1  6            103

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
	ds.BSTInsert(root, 20)

	fmt.Println(ds.BSTSearch(root, 18))
	fmt.Println(ds.BSTSearch(root, 3))
	fmt.Println(ds.BSTSearch(root, 21))

	ds.BSTDelete(root, 12)

	fmt.Println(root.Data)
	fmt.Println(root.Left.Data)
	fmt.Println(root.Right.Left)

	heap := ds.NewHeap(arr[0])
	for _, val := range arr[1:] {
		heap.Insert(val)
	}

	heap.Print()
	fmt.Println("===")

	heap.Delete()
	heap.Delete()
	heap.Delete()

	heap.Print()

	fmt.Println("===")

	t := ds.NewTrie()
	t.Insert("cat")
	t.Insert("cab")
	t.Insert("call")
	t.Insert("caller")

	fmt.Println(t.Root.Kids["c"].Kids["a"].Kids)

	t.CollectAllWords(nil, "")

	for _, word := range t.Words {
		fmt.Println(word)
	}

	fmt.Println("===")

	// create graph, vertices, and edges
	g := ds.NewGraph()

	vAlice := ds.NewVertex("alice")

	vBob := ds.NewVertex("bob")
	vCandy := ds.NewVertex("candy")
	vDerek := ds.NewVertex("derek")
	vElaine := ds.NewVertex("elaine")

	vFred := ds.NewVertex("fred")
	vGina := ds.NewVertex("gina")

	vHelen := ds.NewVertex("helen")
	vIrena := ds.NewVertex("irena")

	g.AddVertex(vAlice)

	g.AddVertex(vBob)
	g.AddVertex(vCandy)
	g.AddVertex(vDerek)
	g.AddVertex(vElaine)

	g.AddVertex(vFred)
	g.AddVertex(vGina)

	g.AddVertex(vHelen)
	g.AddVertex(vIrena)

	vAlice.AddAdjacentVertex(vBob)
	vAlice.AddAdjacentVertex(vCandy)
	vAlice.AddAdjacentVertex(vDerek)
	vAlice.AddAdjacentVertex(vElaine)

	vBob.AddAdjacentVertex(vCandy)
	vBob.AddAdjacentVertex(vFred)

	vCandy.AddAdjacentVertex(vAlice)
	vCandy.AddAdjacentVertex(vHelen)

	vDerek.AddAdjacentVertex(vAlice)
	vDerek.AddAdjacentVertex(vElaine)
	vDerek.AddAdjacentVertex(vGina)

	vElaine.AddAdjacentVertex(vAlice)
	vElaine.AddAdjacentVertex(vDerek)

	vFred.AddAdjacentVertex(vBob)
	vFred.AddAdjacentVertex(vHelen)

	vGina.AddAdjacentVertex(vDerek)
	vGina.AddAdjacentVertex(vIrena)

	vHelen.AddAdjacentVertex(vFred)
	vHelen.AddAdjacentVertex(vCandy)

	vIrena.AddAdjacentVertex(vGina)

	visited := make(map[string]bool)

	fmt.Println("=== dfs ===")
	g.DFSTraverse(vAlice, visited)

	fmt.Println("=== bfs ===")
	g.BFSTraverese(vAlice)

}
