package algorithms

// KadaneMaxSumSubArray finds the maximum sum sub-array in
// the array using Kadane's algorithm.
func KadaneMaxSumSubArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	ans := arr[0]
	sum := 0
	for i := 0; i < len(arr); i++ {
		if sum+arr[i] > 0 {
			sum += arr[i]
		} else {
			sum = 0
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}
