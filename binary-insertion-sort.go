package algorithms

// BinaryInsertionSort performs a binary insertion sort on an
// array of integers.
func BinaryInsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		currentElement := arr[i]
		pos := binaryInsertionSortSearch(arr, currentElement, 0, i)
		j := i
		for j > pos {
			arr[j] = arr[j-1]
			j--
		}
		arr[pos] = currentElement
	}
}

// binaryInsertionSortSearch is a helper function for performing a binary search
// on the sorted part of the integer array while performing a binary insertion
// sort.
func binaryInsertionSortSearch(arr []int, item, start, end int) int {
	for start < end {
		mid := (start + end) / 2
		if arr[mid] <= item {
			start = mid + 1
			continue
		}
		end = mid
	}
	return start
}
