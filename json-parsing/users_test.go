package jsonparsing

import (
	"testing"

	"gotest.tools/assert"
)

func TestUsers(t *testing.T) {
	t.Run("ShouldBeAbleToReturnEncodedJsonOfUserData", func(t *testing.T) {
		user := User{Name: "Ironman", Email: "ironman@gmail.com", Passsword: "abc", IsActive: true}

		want := `{"name":"Ironman","email":"ironman@gmail.com","isActive":true}`

		got := EncodeJson(user)

		assert.Equal(t, want, got)
	})

	t.Run("ShouldBeAbleToReturnDecodedJsonOfUserData", func(t *testing.T) {
		userJson := `{"name":"Ironman","email":"ironman@gmail.com","isActive":true}`

		want := &User{Name: "Ironman", Email: "ironman@gmail.com", Passsword: "", IsActive: true}

		got, _ := DecodeJson(userJson)

		assert.DeepEqual(t, want, got)
	})

	t.Run("ShouldBeAbleToReturnErrorWhenWeProvideInvalidJsonOfUserData", func(t *testing.T) {
		userJson := `{"name":"Ironman","email""ironman@gmail.com","isActive":true}`

		want := "invalid json string"

		_, err := DecodeJson(userJson)

		assert.Equal(t, want, err.Error())
	})
}
