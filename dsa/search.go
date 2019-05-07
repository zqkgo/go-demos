package dsa

// 二分查找
func BinarySearch(elements []int, target int) int {
	length := len(elements)
	if length == 0 {
		return -1
	}
	low := 0
	high := length
	for low <= high {
		mid := (low + high) / 2
		if elements[mid] == target {
			return mid
		} else if elements[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 普通查找
func Search(elements []int, target int) int {
	for i := 0; i < len(elements); i++ {
		if elements[i] == target {
			return i
		}
	}
	return -1
}
