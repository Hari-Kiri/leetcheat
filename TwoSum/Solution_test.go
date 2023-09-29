package TwoSum

import (
	"fmt"
	"testing"
)

type solutionTest struct {
	nums     []int
	target   int
	expected []int
}

var solutionTests = []solutionTest{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
	{[]int{0, 4, 3, 0}, 0, []int{0, 3}},
	{[]int{0, 3, -3, 4, -1}, -1, []int{0, 4}},
}

func TestSolution(t *testing.T) {
	for i := 0; i < len(solutionTests); i++ {
		if output := Solution(solutionTests[i].nums, solutionTests[i].target); fmt.Sprint(output) != fmt.Sprint(solutionTests[i].expected) {
			t.Errorf("Output %d not equal to expected %d", output, solutionTests[i].expected)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution([]int{2, 7, 11, 15}, 9)
	}
}

func ExampleSolution() {
	fmt.Println(Solution([]int{2, 7, 11, 15}, 9))
	// Output: [0 1]
}
