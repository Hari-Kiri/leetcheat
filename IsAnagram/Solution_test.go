package IsAnagram

import "testing"

func TestSolution(t *testing.T) {

	got := Solution("anagram", "nagaram")
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
