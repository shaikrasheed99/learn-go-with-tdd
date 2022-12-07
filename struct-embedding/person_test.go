package structembedding

import (
	"testing"

	"gotest.tools/assert"
)

func TestPerson(t *testing.T) {
	p := Person{
		name: "Ironman",
		age:  3000,
		address: Address{
			state:   "Mumbai",
			country: "India",
		},
	}

	t.Run("Should be able to change state of a person", func(t *testing.T) {
		state := "Telangana"
		want := p
		want.address.state = "Telangana"

		got := p.modifyState(state)

		assert.Equal(t, want, got)
	})

	t.Run("Should be able to change age of a person", func(t *testing.T) {
		age := 4000
		want := p
		want.age = 4000

		got := p.modifyAge(age)

		assert.Equal(t, want, got)
	})
}
