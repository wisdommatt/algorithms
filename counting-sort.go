package algorithms

// CountingSort performs a counting sort on an array of integers
// and return the sorted array.
func CountingSort(arr []int, max int) []int {
	positions := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		positions[arr[i]] += 1
	}
	for i := 1; i < len(positions); i++ {
		k := i - 1
		positions[i] = positions[k] + positions[i]
	}
	output := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		k := arr[i]
		pos := positions[k]
		output[pos-1] = k
		positions[k]--
	}
	return output
}
