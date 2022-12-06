package helloworld

import (
	"testing"

	"gotest.tools/assert"
)

func TestHelloWorld(test *testing.T) {
	want := "Hello World"

	got := HelloWorld()

	assert.Equal(test, got, want)
}

func TestGreet(test *testing.T) {
	test.Run("Should be able to say Hello with name", func(t *testing.T) {
		want := "Hello, Ironman"

		got := Greet("Ironman")

		assert.Equal(t, want, got)
	})

	test.Run("Should be able to say Hello when we provide empty name", func(t *testing.T) {
		want := "Hello"

		got := Greet("")

		assert.Equal(t, want, got)
	})
}
