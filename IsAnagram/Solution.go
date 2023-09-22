package IsAnagram

import (
	"fmt"
)

// Check two string is anagram or not.
// The goal is to check count every alphabet in s & t
// then compare it has same count or not.
func Solution(s string, t string) bool {
	// Stop when s & t not in same length
	if len(s) != len(t) {
		return false
	}

	temp0 := make(map[string]int)
	temp1 := make(map[string]int)
	for i := 0; i < len(s); i++ {
		temp0[string(s[i])] = temp0[string(s[i])] + 1
		temp1[string(t[i])] = temp1[string(t[i])] + 1
	}

	return fmt.Sprint(temp0) == fmt.Sprint(temp1)
}
