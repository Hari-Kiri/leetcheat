package TwoSum

/*Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
The solution is to sum each number in array with other number in that array itself to get the target.*/
func Solution(nums []int, target int) []int {
	// Loop through nums array.
	for i := 0; i < len(nums); i++ {

		// Loop againt nums array to sum with i loop.
		for j := 0; j < len(nums); j++ {

			// When loop j not in same index with i loop, sum it. If it same with target return.
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}

		}

	}

	return nil
}
