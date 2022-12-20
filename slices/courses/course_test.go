package courses

import (
	"testing"

	"gotest.tools/assert"
)

func TestCourse(t *testing.T) {
	type TestCase struct {
		description       string
		courses           []string
		courseToBeRemoved string
		want              []string
	}

	testCases := []TestCase{
		{
			description:       "ShouldBeAbleToRemoveSingleOccuranceOfACourseFromCoursesSlice",
			courses:           []string{"Java", "JavaScript", "Python", "C", "C++"},
			courseToBeRemoved: "Python",
			want:              []string{"Java", "JavaScript", "C", "C++"},
		},
		{
			description:       "ShouldBeAbleToRemoveAllTheOccurancesOfACourseFromCoursesSlice",
			courses:           []string{"Java", "C", "JavaScript", "Python", "C", "C++", "C"},
			courseToBeRemoved: "C",
			want:              []string{"Java", "JavaScript", "Python", "C++"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			got := Remove(testCase.courses, testCase.courseToBeRemoved)

			assert.DeepEqual(t, testCase.want, got)
		})
	}
}
