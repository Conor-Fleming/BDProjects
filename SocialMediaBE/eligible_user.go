package main

import (
	"errors"
)

func userIsEligible(email, password string, age int) error {
	if email == "" {
		return errors.New("email can't be empty")
	}
	if password == "" {
		return errors.New("password can't be empty")
	}
	if age < 18 {
		return errors.New("age must be at least 18 years old")
	}
	return nil
}
