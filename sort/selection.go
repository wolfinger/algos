package sort

func Selection(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		if i != min {
			temp := arr[i]
			arr[i] = arr[min]
			arr[min] = temp
		}
	}

	return arr
}
