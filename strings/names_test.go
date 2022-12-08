package strings

import (
	"testing"

	"gotest.tools/assert"
)

func TestShouldBeAbleToReturnUniqueNamesFromTwoGroupOfNames(t *testing.T) {
	groupOne := []string{"Ava", "Emma", "Olivia"}
	groupTwo := []string{"Olivia", "Sophia", "Emma"}
	want := []string{"Ava", "Emma", "Olivia", "Sophia"}

	got := UniqueNames(groupOne, groupTwo)

	assert.DeepEqual(t, want, got)
}

func TestUniqueCharacters(t *testing.T) {
	t.Run("Should be able to return number of unique characters of two names", func(t *testing.T) {
		firstName := "Jack"
		secondName := "Mack"
		want := 2

		got, error := UniqueCharacters(firstName, secondName)

		if error != nil {
			t.Error()
		}

		assert.Equal(t, want, got)
	})

	t.Run("Should be able to return error when the length two names are not equal", func(t *testing.T) {
		firstName := "Thanos"
		secondName := "Thor"
		want := "both names did not have same length"

		_, error := UniqueCharacters(firstName, secondName)

		if error == nil {
			t.Error()
		}

		assert.Equal(t, want, error.Error())
	})
}
