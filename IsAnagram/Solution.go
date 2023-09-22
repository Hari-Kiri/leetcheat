package IsAnagram

import (
	"strings"
)

// Check two string is anagram or not.
// The goal is to check count every alphabet in s & t
// then compare it has same count or not.
func Solution(s string, t string) bool {
	// Stop when s & t not in same length
	if len(s) != len(t) {
		return false
	}

	// Lower the string
	sLower := strings.ToLower(s)
	tLower := strings.ToLower(t)
	// Create data map of integer or any othe key = value pair data type in other language
	temp := make(map[string]int)
	// Loop through sLower
	for i := 0; i < len(sLower); i++ {
		// Add new key to temp using sLower[i] as key and increment value
		temp[string(sLower[i])] = temp[string(sLower[i])] + 1
	}

	// Loop through tLower
	for i := 0; i < len(tLower); i++ {
		// Decrease the value in every key.
		// We use tLower[i] as key now and when the key from tLower[i] not present in temp,
		// The compiler will add new key with value minus 1 (-1).
		temp[string(tLower[i])] = temp[string(tLower[i])] - 1
		// Detect -1 then return false
		// When minus 1 detected s & t is not anagram
		if temp[string(tLower[i])] < 0 {
			return false
		}
	}

	return true
}
