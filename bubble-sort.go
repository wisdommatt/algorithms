package algorithms

// BubbleSort performs a bubble sort on an array of
// integers.
//
// Best Case Complexity: 	O(n)
// Average Case Complexity: O(n^2)
// Worst Case Complexity: 	O(n^2)
func BubbleSort(arr []int) {
	j := len(arr) - 1
	for j > 0 {
		swapped := false
		for i := 0; i < j; i++ {
			if arr[i] > arr[i+1] {
				iValue := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = iValue
				swapped = true
			}
		}
		// if there was no swaps it means that the array
		// is already sorted.
		if !swapped {
			break
		}
		j--
	}
}
