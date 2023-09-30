package main

import (
	"fmt"

	"github.com/Hari-Kiri/leetcheat/ConcatenationOfArray"
	"github.com/Hari-Kiri/leetcheat/ContainsDuplicate"
	"github.com/Hari-Kiri/leetcheat/IsAnagram"
	"github.com/Hari-Kiri/leetcheat/TwoSum"
)

func main() {
	fmt.Println(ContainsDuplicate.Solution([]int{22, 14, 2, 18, 22}))
	fmt.Println(IsAnagram.Solution("anagram", "nagaram"))
	fmt.Println(TwoSum.Solution([]int{2, 7, 11, 15}, 18))
	fmt.Println(ConcatenationOfArray.Solution([]int{1, 2, 1}))
}
