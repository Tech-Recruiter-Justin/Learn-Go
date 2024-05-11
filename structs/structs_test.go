package structs_test

import (
	"github.com/stretchr/testify/assert"
	"learnGoWithTests/structs"
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name     string
		shape    structs.Shape
		expected float64
	}{
		{name: "rectangle perimeter", shape: structs.Rectangle{Width: 11.9, Height: 20.1}, expected: 64.0},
		{name: "circle perimeter", shape: structs.Circle{Radius: 10}, expected: 20 * math.Pi},
		{name: "triangle perimeter", shape: structs.Triangle{Base: 10, Left: 10, Right: 10}, expected: 30.0},
	}

	for _, testCase := range perimeterTests {
		assert.Equal(t, testCase.expected, testCase.shape.Perimeter())
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape    structs.Shape
		expected float64
	}{
		{shape: structs.Rectangle{Width: 19, Height: 10}, expected: 190.0},
		{shape: structs.Circle{Radius: 10}, expected: 314.1592653589793},
		{shape: structs.Triangle{Base: 12, Left: 12, Right: 12}, expected: 62.353829072479584},
	}

	for _, testCase := range areaTests {
		assert.Equal(t, testCase.shape.Area(), testCase.expected)
	}
}
