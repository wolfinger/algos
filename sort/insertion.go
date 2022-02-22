package sort

func Insertion(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		temp_val := arr[i]
		j := i - 1

		for j >= 0 {
			if arr[j] > temp_val {
				arr[j+1] = arr[j]
				j -= 1
			} else {
				break
			}
		}

		arr[j+1] = temp_val
	}

	return arr
}
