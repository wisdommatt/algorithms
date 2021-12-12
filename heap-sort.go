package algorithms

// HeapSort performs a heap sort on an array of integers
// and return the sorted array.
func HeapSort(arr []int) []int {
	arrLen := len(arr)
	for i := 1; i < arrLen; i++ {
		heapSortBubbleUp(arr, i)
	}
	sortedArr := []int{}
	length := arrLen - 1
	for i := 0; i < arrLen; i++ {
		sortedArr = append(sortedArr, arr[0])
		if length == 0 {
			continue
		}
		lastVal := arr[length]
		arr = arr[:length]
		arr[0] = lastVal
		heapSortBubbleDown(arr, 0)
		length--
	}
	return sortedArr
}

// heapSortBubbleUp is a helper function to bubble up
// items in a heap sort to maintain the heap property.
func heapSortBubbleUp(arr []int, index int) {
	parent := index / 2
	if arr[parent] <= arr[index] {
		return
	}
	heapSortSwap(arr, index, parent)
	heapSortBubbleUp(arr, parent)
}

// heapSortBubbleDown is a helper function to bubble down
// items in a heap sort to maintain the heap property.
func heapSortBubbleDown(arr []int, index int) {
	leftChildIndex := (2 * index) + 1
	rightChildIndex := (2 * index) + 2

	// checking if right child is < left child; and < arr[index]
	if len(arr) > rightChildIndex && arr[rightChildIndex] < arr[leftChildIndex] {
		if arr[rightChildIndex] < arr[index] {
			heapSortSwap(arr, index, rightChildIndex)
			heapSortBubbleDown(arr, rightChildIndex)
			return
		}
	}

	if len(arr) > leftChildIndex && arr[leftChildIndex] < arr[index] {
		heapSortSwap(arr, index, leftChildIndex)
		heapSortBubbleDown(arr, leftChildIndex)
	}
}

// heapSortSwap is a helper function for swapping array elements
// during a heap sort.
func heapSortSwap(arr []int, index1, index2 int) {
	index1Value := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = index1Value
}
