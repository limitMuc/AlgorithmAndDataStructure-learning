package quickSort

// 快速排序
func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}

// 更容易理解的版本
func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		splitdata := arr[0]
		var less []int
		var greater []int
		var equal []int
		for _, value := range arr[1:] {
			if value < splitdata {
				less = append(less, value)
			} else if value > splitdata {
				greater = append(greater, value)
			} else {
				equal = append(equal, value)
			}
		}

		var result []int
		result = append(result, quicksort(less)...)
		result = append(result, equal...)
		result = append(result, quicksort(greater)...)
		return result
	}
}
