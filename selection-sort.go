package algorithms

// SelectionSort performs a selection sort on an
// array of integers.
//
// Best Case Time Complexity: O(n^2)
// Worst Case Time Complexity: O(n^2)
func SelectionSort(arr []int) {
	j := 0
	arrLen := len(arr)
	for j <= arrLen-1 {
		for i := j + 1; i < arrLen; i++ {
			if arr[i] < arr[j] {
				jValue := arr[j]
				arr[j] = arr[i]
				arr[i] = jValue
			}
		}
		j++
	}
}
