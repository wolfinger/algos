package sort

type sortableArray struct {
	arr []int
}

func Quicksort(arr []int) []int {
	var sortedArr sortableArray

	sortedArr.arr = arr
	sortedArr.quicksortHelper(0, len(sortedArr.arr)-1)

	return sortedArr.arr
}

func (s *sortableArray) quicksortHelper(left int, right int) {
	if (right - left) <= 0 {
		return
	}

	pivot := s.partition(left, right)

	s.quicksortHelper(left, pivot-1)

	s.quicksortHelper(pivot+1, right)
}

func (s *sortableArray) partition(left int, right int) int {
	start := left
	pivot := right
	right -= 1

	for {
		for s.arr[left] < s.arr[pivot] {
			left += 1
		}
		for (right >= start) && (s.arr[right] > s.arr[pivot]) {
			right -= 1
		}

		if left < right {
			s.arr[left], s.arr[right] = s.arr[right], s.arr[left]
			left += 1
		} else {
			break
		}
	}

	s.arr[pivot], s.arr[left] = s.arr[left], s.arr[pivot]

	return left
}
