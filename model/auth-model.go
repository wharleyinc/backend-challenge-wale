package model

import (
	"net/mail"
	"time"
)

type Account struct {
	AccountID string
	Email     string
	CreatedAt time.Time
}

type CreateAccount struct {
	Email        string
	PasswordHash string
	Salt         string
	CreatedAt    time.Time
}

type NewAccount struct {
	AccountID string
	Email     string
	Password  string
}

func (ep CreateAccount) ValidEmail() bool {
	_, err := mail.ParseAddress(ep.Email)
	return err == nil
}

func (ep CreateAccount) ValidPassword() bool {
	if !(len(ep.PasswordHash) >= 6 && isAlphanumeric(ep.PasswordHash)) {
		return false
	}

	return true
}

func isAlphanumeric(password string) bool {
	hasNumber := false
	hasAlphabet := false
	for _, c := range password {
		if isAlphabet(c) {
			hasAlphabet = true
		} else if isNumber(c) {
			hasNumber = true
		} else {
			// We only support alphanumeric.
			return false
		}
	}

	return hasNumber && hasAlphabet
}

func isAlphabet(c rune) bool {
	for i := 'A'; i <= 'Z'; i++ {
		// Checks lower and upper case.
		if string(i) == string(c) || string(i+32) == string(c) {
			return true
		}
	}

	return false
}

func isNumber(c rune) bool {
	for i := '0'; i <= '9'; i++ {
		if string(i) == string(c) {
			return true
		}
	}

	return false
}

type AuthenticateAccount struct {
	Email    string
	Password string
}
