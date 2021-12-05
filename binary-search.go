package algorithms

// BinarySearch performs a binary search on the array to
// find the position of the given element.
//
// NOTE: binary search should only be used with a sorted array.
// Using binary search with an unsorted array will return an incorrect
// result.
func BinarySearch(arr []int, val int) int {
	arrLen := len(arr)
	return binarySearchRecursive(arr, val, 0, arrLen)
}

// binarySearchRecursive is a helper function to perform a recursive
// binary search on an array.
func binarySearchRecursive(arr []int, val, left, right int) int {
	mid := (left + right) / 2
	if mid >= len(arr) || right < 0 {
		return -1
	}
	if arr[mid] == val {
		return mid
	}
	if val < arr[mid] {
		return binarySearchRecursive(arr, val, left, mid-1)
	}
	return binarySearchRecursive(arr, val, mid+1, right)
}
