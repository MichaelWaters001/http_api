package main

import (
	"errors"
)

type user struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ensure all user fields are populated
func (newUser user) ValidateNewUser() error {
	if newUser.ID == "" {
		return errors.New("ID not set")
	}
	if newUser.Email == "" {
		return errors.New("Email not set")
	}
	if newUser.Name == "" {
		return errors.New("Name not set")
	}
	return nil
}

// mock user DB
var users = []user{
	{ID: "1", Name: "Foo", Email: "foo@example.com"},
	{ID: "2", Name: "Bar", Email: "Bar@example.com"},
	{ID: "3", Name: "John", Email: "John@Doh.com"},
}

type email struct {
	Email string `json:"email"`
}
