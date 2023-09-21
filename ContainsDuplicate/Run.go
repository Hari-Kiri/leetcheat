package ContainsDuplicate

// Check if any duplicate in array.
func Run(nums []int) bool {
	// Declare return value variable
	var anyDuplicate bool

	// Declare data map or other key = value pair data in other language
	temp := make(map[int]bool, 0)

	// Loop through nums
	for i := 0; i < len(nums); i++ {

		// Check if any key named with nums[i] with value true.
		// If yes set return value to true then break the loop
		if temp[nums[i]] {
			anyDuplicate = true
			break
		}

		// If not any key named with nums[i] with value true,
		// set it for next loop
		temp[nums[i]] = true

	}

	return anyDuplicate
}
