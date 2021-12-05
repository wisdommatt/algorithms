package algorithms

// MergeSort performs a merge sort on an array of integers
// and return the sorted array.
func MergeSort(arr []int) []int {
	arrLen := len(arr)
	mid := arrLen / 2
	if arrLen <= 1 {
		return arr
	}
	a := MergeSort(arr[:mid])
	b := MergeSort(arr[mid:])
	return mergeSortMergeArray(a, b)
}

// mergeSortMergeArray is a helper function for merging/sorting
// two arrays during a merge sort.
func mergeSortMergeArray(arr1 []int, arr2 []int) []int {
	newArr := []int{}
	i := 0
	j := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			newArr = append(newArr, arr1[i])
			i++
			continue
		}
		newArr = append(newArr, arr2[j])
		j++
	}
	if i < len(arr1) {
		newArr = append(newArr, arr1[i:]...)
	}
	if j < len(arr2) {
		newArr = append(newArr, arr2[j:]...)
	}
	return newArr
}
