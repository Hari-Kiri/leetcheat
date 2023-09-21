package IsAnagram

import (
	"strings"
)

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
		// Add new key to temp using sLower[i] as key and value integer + 1 using ++ (plus plus) operator
		temp[string(sLower[i])]++
	}

	// Loop through sLower
	for i := 0; i < len(tLower); i++ {
		// Decrease the value in every key using -- (minus minus) operator.
		// We use tLower[i] as key now and when the key from tLower[i] not present in temp,
		// The compiler will add new key with value minus 1 (-1).
		temp[string(tLower[i])]--
		// Detect -1 then return false
		// When minus 1 detected s & t is not anagram
		if temp[string(tLower[i])] < 0 {
			return false
		}
	}

	return true
}
