package algorithms

// InsertionSort performs an insertion sort on a
// slice of integers.
func InsertionSort(arr []int) {
	arrLen := len(arr)
	for j := 1; j < arrLen; j++ {
		i := j - 1
		for i >= 0 && arr[i] > arr[i+1] {
			jValue := arr[i+1]
			arr[i+1] = arr[i]
			arr[i] = jValue
			i--
		}
	}
}
