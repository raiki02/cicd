package tests

import "testing"

func TestAdd(t *testing.T) {
	type testCase struct {
		a, b int
	}

	a := 1
	b := 2
	want := a + b

	got := Add(a, b)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}
