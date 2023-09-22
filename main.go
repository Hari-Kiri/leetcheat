package main

import (
	"fmt"

	"github.com/Hari-Kiri/leetcheat/ContainsDuplicate"
	"github.com/Hari-Kiri/leetcheat/IsAnagram"
)

func main() {
	fmt.Println(ContainsDuplicate.Solution([]int{22, 14, 2, 18, 22}))
	fmt.Println(IsAnagram.Solution("anagram", "nagaram"))
}
