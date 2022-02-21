package main

import (
	"fmt"

	"github.com/wolfinger/algos/search"
)

func main() {
	// test binary search
	arr := []int{1, 3, 6, 7, 8, 12, 18, 24, 52, 103}

	idx, found := search.Binary(arr, 6)
	fmt.Println(idx, found)

	idx, found = search.Binary(arr, 52)
	fmt.Println(idx, found)

	idx, found = search.Binary(arr, 1)
	fmt.Println(idx, found)

	idx, found = search.Binary(arr, 2)
	fmt.Println(idx, found)
}
