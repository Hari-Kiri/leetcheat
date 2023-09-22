package IsAnagram

import (
	"fmt"
	"testing"
)

type solutionTest struct {
	s, t     string
	expected bool
}

var solutionTests = []solutionTest{
	{"anagram", "nagaram", true},
	{"rat", "car", false},
	{"ac", "bb", false},
	{"aacc", "ccac", false},
	{"caca", "ccaa", true},
	{"qwertyuiopasdfghjklzxcvbnm", "lkjhgfdsaqwertyuiopmnbvcxz", true},
	{"poiuqwertyhjklgfdsabnmvcxz", "qazwsxedcrfvtgbyhnujmikol", false},
}

func TestSolution(t *testing.T) {
	for i := 0; i < len(solutionTests); i++ {
		if output := Solution(solutionTests[i].s, solutionTests[i].t); output != solutionTests[i].expected {
			t.Errorf("Output %t not equal to expected %t", output, solutionTests[i].expected)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution("mnbvcxzlkjhgfdsaqwertyuiop", "ghjklfdsapqowieurytbvncmxz")
	}
}

func ExampleSolution() {
	fmt.Println(Solution("polisi", "isilop"))
	// Output: true
}
