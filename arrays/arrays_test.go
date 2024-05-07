package arrays_test

import (
	"learnGoWithTests/arrays"
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	actual := arrays.Sum(nums)
	expected := 15
	assertEquals(t, actual, expected)
}

func assertEquals(t *testing.T, actual, expected int) {
	if actual != expected {
		t.Errorf("actual: %d, expected: %d", actual, expected)
	}
}
