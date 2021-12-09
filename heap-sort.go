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
	indexVal := arr[index]
	if arr[parent] <= indexVal {
		return
	}
	arr[index] = arr[parent]
	arr[parent] = indexVal
	heapSortBubbleUp(arr, parent)
}

// heapSortBubbleDown is a helper function to bubble down
// items in a heap sort to maintain the heap property.
func heapSortBubbleDown(arr []int, index int) {
	leftChildIndex := (2 * index) + 1
	rightChildIndex := (2 * index) + 2
	indexVal := arr[index]

	// checking if right child is < left child; and < arr[index]
	if len(arr) > rightChildIndex && arr[rightChildIndex] < arr[leftChildIndex] {
		if arr[rightChildIndex] < indexVal {
			arr[index] = arr[rightChildIndex]
			arr[rightChildIndex] = indexVal
			heapSortBubbleDown(arr, rightChildIndex)
			return
		}
	}

	if len(arr) > leftChildIndex && arr[leftChildIndex] < indexVal {
		arr[index] = arr[leftChildIndex]
		arr[leftChildIndex] = indexVal
		heapSortBubbleDown(arr, leftChildIndex)
	}
}
