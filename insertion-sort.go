package algorithms

// InsertionSort performs an insertion sort on a
// slice of integers.
func InsertionSort(arr []int) {
	arrLen := len(arr)
	for i := 1; i < arrLen; i++ {
		j := i - 1
		insertionSortSwap(arr, j, i)
	}
}

// insertionSortSwap is a helper function for swapping
// values in an array during an insertion sort.
func insertionSortSwap(arr []int, j, i int) {
	if (j < 0 || i < 0) || arr[i] >= arr[j] {
		return
	}
	iValue := arr[i]
	arr[i] = arr[j]
	arr[j] = iValue
	insertionSortSwap(arr, j-1, i-1)
}
