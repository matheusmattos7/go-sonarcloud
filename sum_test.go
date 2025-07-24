package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result != 5 {
		t.Error("The result must be 5")
	}
}

func TestSub(t *testing.T) {
	got := sub(5, 3)
	want := 2

	if got != want {
		t.Errorf("sub(5, 3) = %d; want %d", got, want)
	}
}

func TestTimes(t *testing.T) {
	got := times(4, 3)
	want := 12

	if got != want {
		t.Errorf("times(4, 3) = %d; want %d", got, want)
	}
}

func TestSumX(t *testing.T) {
	got := sumX(2, 3) // 2 + 3 + 2 = 7
	want := 7

	if got != want {
		t.Errorf("sumX(2, 3) = %d; want %d", got, want)
	}
}
