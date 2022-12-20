package jsonparsing

import (
	"encoding/json"
	"errors"
)

var ErrInvalidJson error = errors.New("invalid json string")

type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Passsword string `json:"-"`
	IsActive  bool   `json:"isActive"`
}

func EncodeJson(user User) string {
	value, _ := json.Marshal(user)

	return string(value)
}

func DecodeJson(userJsonString string) (*User, error) {
	userJsonBytes := []byte(userJsonString)

	isValid := json.Valid(userJsonBytes)

	if isValid {
		var user *User
		json.Unmarshal(userJsonBytes, &user)
		return user, nil
	} else {
		return nil, ErrInvalidJson
	}
}
