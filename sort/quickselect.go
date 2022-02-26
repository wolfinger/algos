package sort

func Quickselect(arr []int, k int) int {
	var sortedArr sortableArray

	sortedArr.arr = arr
	kthLargest := sortedArr.quickselectHelper(0, len(sortedArr.arr)-1, k)

	return kthLargest
}

func (s *sortableArray) quickselectHelper(left int, right int, k int) int {
	if (right - left) <= 0 {
		return s.arr[left]
	}

	pivot := s.partition(left, right)

	if k < pivot {
		return s.quickselectHelper(left, pivot-1, k)
	} else if k > pivot {
		return s.quickselectHelper(pivot+1, right, k)
	} else {
		return s.arr[pivot]
	}
}
