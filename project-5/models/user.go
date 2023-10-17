package models

import "errors"

var Users = map[uint64]User{
	123: {
		FName:    "Bob",
		LName:    "abc",
		Password: "someSecretPassword",
		Email:    "bob@email.com",
	},
}

func FetchUser(userId uint64) (User, error) {

	v, ok := Users[userId]

	if !ok {
		return User{}, errors.New("user not found")
	}
	return v, nil
}
