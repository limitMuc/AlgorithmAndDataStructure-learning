package selectSort

// 选择排序
func selectSort(arr []int) {
	var minindex int
	for i := 0; i < len(arr)-1; i++ {
		minindex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minindex] {
				minindex = j
			}
		}

		arr[i], arr[minindex] = arr[minindex], arr[i]
	}
}