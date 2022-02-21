package search

func Binary(searchArr []int, searchVal int) (int, bool) {
	var low, mid, high int
	arrLen := len(searchArr)

	if arrLen == 1 {
		if searchArr[0] == searchVal {
			return 0, true
		} else {
			return 0, false
		}
	}

	low = 0
	high = arrLen - 1

	for low <= high {
		mid = (high + low) / 2
		if searchArr[mid] == searchVal {
			return mid, true
		} else if searchVal < searchArr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if searchArr[high] == searchVal {
		return high, true
	} else if searchArr[low] == searchVal {
		return low, true
	}

	return 0, false
}
