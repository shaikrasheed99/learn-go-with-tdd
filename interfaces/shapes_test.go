package interfaces

import (
	"testing"

	"gotest.tools/assert"
)

func TestShapes(t *testing.T) {
	type TestCase struct {
		name           string
		shape          Shape
		wantedArea     float64
		wantedPerimter float64
	}

	testCases := []TestCase{
		{
			name:           "Rectangle",
			shape:          &Rectangle{length: 10, breadth: 20},
			wantedArea:     200.0,
			wantedPerimter: 60.0,
		},
		{
			name:           "Square",
			shape:          &Square{side: 10},
			wantedArea:     100.0,
			wantedPerimter: 40.0,
		},
		{
			name:           "Circle",
			shape:          &Circle{radius: 10},
			wantedArea:     314.0,
			wantedPerimter: 62.800000000000004,
		},
	}

	for _, testCase := range testCases {
		t.Run("Area of "+testCase.name, func(t *testing.T) {
			got := testCase.shape.area()

			assert.Equal(t, testCase.wantedArea, got)
		})
	}

	for _, testCase := range testCases {
		t.Run("Perimeter of "+testCase.name, func(t *testing.T) {
			got := testCase.shape.perimeter()

			assert.Equal(t, testCase.wantedPerimter, got)
		})
	}
}
