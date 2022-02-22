package sort

func Bubble(arr []int) []int {
	sorted := false
	maxIdx := len(arr) - 1

	for !sorted {
		sorted = true
		for i := 0; i < maxIdx; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp

				sorted = false
			}
		}
		maxIdx -= 1
	}

	return arr
}
