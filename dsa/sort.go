package dsa

// 插入排序，时间复杂度O(n^2)
func InsertionSort(elements []int) {
	length := len(elements)
	if length < 2 {
		return
	}
	for i := 1; i < length; i++ {
		cur := elements[i]
		j := i - 1
		// 从前一位开始，比cur大的元素后移
		// 这里是逐个比较，因已经有序，故可使用二分查找提升效率
		for j >= 0 && elements[j] > cur {
			elements[j+1] = elements[j]
			j--
		}
		elements[j+1] = cur
	}
}

// 冒泡排序，时间复杂度O(n^2)
func BubbleSort(elements []int) {
	length := len(elements)
	if length < 2 {
		return
	}
	end := length - 1 // 标记最后一位，不断向前移动
	swapped := true   // 标识是否发生交换
	for swapped {
		swapped = false
		start := 0 // 每次都从第一位开始逐个向后比较
		for start < end {
			if elements[start] > elements[start+1] {
				elements[start], elements[start+1] = elements[start+1], elements[start]
				swapped = true
			}
			start++
		}
		end--
	}
}

// 快速排序，平均时间复杂度O(nlogn)，最坏O(n^2)
// TODO: 不建新数组
func QuickSort(elements []int) []int {
	length := len(elements)
	if length < 2 {
		return elements
	}
	pivot := elements[length/2]
	var left, right, pivots []int
	for i := 0; i < length; i++ {
		if elements[i] > pivot {
			right = append(right, elements[i])
		} else if elements[i] < pivot {
			left = append(left, elements[i])
		} else {
			pivots = append(pivots, elements[i])
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)
	var newElements []int
	newElements = append(newElements, left...)
	newElements = append(newElements, pivots...)
	newElements = append(newElements, right...)
	return newElements
}

// 选择排序，时间复杂度O(n^2)
func SelectionSort(elements []int) {
	length := len(elements)
	if length < 2 {
		return
	}
	for i := 0; i < length; i++ {
		minKey := i
		for j := i + 1; j < length; j++ { // 每次选择后面值最小的元素与i位置元素交换
			if elements[minKey] > elements[j] {
				minKey = j
			}
		}
		elements[i], elements[minKey] = elements[minKey], elements[i]
	}
}

// 归并排序，时间复杂度O(nlogn)
func MergeSort(elements []int) []int {
	length := len(elements)
	if length < 2 {
		return elements
	}
	mid := length / 2
	left := elements[:mid]
	right := elements[mid:]
	left = MergeSort(left)
	right = MergeSort(right)
	merged := merge(left, right)
	return merged
}

func merge(left ,right []int) []int {
	var newEles []int
	var i, j int
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			newEles = append(newEles, left[i])
			i++
		} else {
			newEles = append(newEles, right[j])
			j++
		}
	}
	for i < len(left) {
		newEles = append(newEles, left[i])
		i++
	}
	for j < len(right) {
		newEles = append(newEles, right[j])
		j++
	}
	return newEles
}
