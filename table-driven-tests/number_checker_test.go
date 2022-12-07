package tabledriventests

import (
	"testing"

	"gotest.tools/assert"
)

func TestNumberChecker(t *testing.T) {
	type TestCase struct {
		name   string
		number int
		want   string
	}

	testCases := []TestCase{
		{
			name:   "Should be able to check number as positive",
			number: 10,
			want:   "Positive",
		},
		{
			name:   "Should be able to check number as negative",
			number: -10,
			want:   "Negative",
		},
		{
			name:   "Should be able to check number as zero",
			number: 0,
			want:   "Zero",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := WhichNumber(testCase.number)

			assert.Equal(t, testCase.want, got)
		})
	}
}
