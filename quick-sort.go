package algorithms

// QuickSort performs a quick sort on an array of
// integers.
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSortRecursive(arr, 0, len(arr)-1)
}

// quickSortRecursive is a helper function to recursively
// perform a quick sort on an array of integers.
func quickSortRecursive(arr []int, lo, hi int) {
	if lo < hi {
		p := quickSortPartition(arr, lo, hi)
		quickSortRecursive(arr, lo, p-1)
		quickSortRecursive(arr, p+1, hi)
	}
}

// quickSortPartition is a helper function to partition an
// array of integers during a quick sort and return the next
// partition index.
func quickSortPartition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[hi] = arr[hi], arr[i]
	return i
}
