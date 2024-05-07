package integers_test

import (
	"fmt"
	"learnGoWithTests/integers"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := integers.Add(2, 2)
	expected := 4
	assertEquals(t, sum, expected)
}

func ExampleAdd() {
	sum := integers.Add(1, 2)
	fmt.Println(sum)
	// Output: 3
}

func assertEquals(t *testing.T, actual, expected int) {
	if actual != expected {
		t.Errorf("actual: %d, expected: %d", actual, expected)
	}
}
