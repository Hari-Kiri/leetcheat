package IsAnagram

// Check two string is anagram or not.
// The goal is to count up every letter in s
// then count it down with t. If there is a
// letter in t which is not in s the count down
// will set it as -1, so that is we detect not
// anagram.
func Solution(s string, t string) bool {
	// Stop when s & t not in same length.
	// Length not same is absolutely not
	// anagram.
	if len(s) != len(t) {
		return false
	}

	// Create a temporary array variable
	// with length of 26, which is representing
	// 26 letters of alphabet
	temp := make([]int, 26)
	// Loop through s. For loop in go will
	// convert the string value to byte.
	// We will reduce the byte value each letter in s
	// with 97. FYI 97 is byte of small letter 'a'.
	// a = 97, b = 98, c = 99 ...
	for i := 0; i < len(s); i++ {
		// Byte of the s[i] value - 97 will get
		// the index of temporary array and
		// increment by 1 the value in that index
		index := int(s[i] - 97)
		temp[index] += 1
	}
	// Loop through t and count down every alphabet count of s
	// inside temporary slice variable.
	for i := 0; i < len(t); i++ {
		index := int(t[i] - 97)
		temp[index] -= 1
		// detect the -1 value then we get it not anagram
		if temp[index] < 0 {
			return false
		}
	}

	return true
}
