package iteration_test

import (
	"learnGoWithTests/iteration"
	"testing"
)

func TestRepeat(t *testing.T) {
	actual := iteration.Repeat("a", 7)
	expected := "aaaaaaa"
	assertEquals(t, actual, expected)
}

func BenchmarkRepeat(b *testing.B) {
	iteration.Repeat("a", 7)
}

func assertEquals(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Errorf("actual: %q, expected: %q", actual, expected)
	}
}
