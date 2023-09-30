package ConcatenationOfArray

func Solution(nums []int) []int {
	return append(nums, nums...)
}
