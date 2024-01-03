package main

// 1480. Running Sum of 1d Array
func RunningSum(nums []int) []int {
	sum := []int{nums[0]}
	for i := 0; i < len(nums)-1; i++ {
		add := sum[i] + nums[i+1]
		sum = append(sum, add)
	}
	return sum
}
