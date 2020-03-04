package insertSort

// 插入排序
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		current := arr[i]
		preindex := i - 1
		for preindex >= 0 && arr[preindex] > current {
			arr[preindex+1] = arr[preindex]
			preindex--
		}
		arr[preindex+1] = current
	}
}