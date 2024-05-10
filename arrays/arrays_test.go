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

func TestSumAll(t *testing.T) {
	nums := [][]int{{1, 2}, {3, 4}, {5, 6}}
	actual := arrays.SumAll(nums, false)
	expected := []int{3, 7, 11}
	assertArraysEqual(t, actual, expected)
}

func TestSumAllTails(t *testing.T) {
	nums := [][]int{{1, 2, 3}, {3, 4, 9}, {5, 6}}
	actual := arrays.SumAll(nums, true)
	expected := []int{5, 13, 6}
	assertArraysEqual(t, actual, expected)
}

func assertEquals(t *testing.T, actual, expected int) {
	if actual != expected {
		t.Errorf("actual: %d, expected: %d", actual, expected)
	}
}

func assertArraysEqual(t *testing.T, actual, expected []int) {
	if len(actual) != len(expected) {
		t.Errorf("arrays are not with equal length, actualLength: %d, expectedLength: %d", len(actual), len(expected))
	}
	for i, n := range actual {
		if n != expected[i] {
			t.Errorf("arrays are not equal, actual: %v, expected: %v", actual, expected)
		}
	}
}
