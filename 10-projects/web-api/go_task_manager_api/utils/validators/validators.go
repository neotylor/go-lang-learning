package validators

import (
	"errors"
	"unicode"
)

func ValidateUserInput(username, password string) error {
	if len(username) < 3 {
		return errors.New("username must be at least 3 characters")
	}
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	if !containsDigit(password) {
		return errors.New("password must include at least one number")
	}
	if !containsSymbol(password) {
		return errors.New("password must include at least one special character")
	}
	return nil
}

func containsDigit(s string) bool {
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			return true
		}
	}
	return false
}

func containsSymbol(s string) bool {
	for _, ch := range s {
		if unicode.IsSymbol(ch) || unicode.IsPunct(ch) {
			return true
		}
	}
	return false
}
