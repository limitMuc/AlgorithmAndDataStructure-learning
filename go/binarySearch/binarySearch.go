package binarySearch

// 二分查找
func binarySearch(a []int, lookfor int) int {
	var low int = 0
	var height int = len(a) - 1
	for low <= height {
		var mid int = low + (height-low)/2
		if a[mid] == lookfor {
			return mid
		} else if a[mid] > lookfor {
			height = mid - 1
		} else if a[mid] < lookfor {
			low = mid + 1
		}

	}

	return -1
}
