package model

import (
	"helpdesk-ticketing-security/errors"
	"net/mail"
)

type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *AuthLoginRequest) Validate() error {
	fields, errors := make(errors.ValidationError), error(nil)

	if a.Email == "" {
		fields["email"] = "email is required"

	} else if _, err := mail.ParseAddress(a.Email); err != nil {
		fields["email"] = "email must be valid"
	}
	if a.Password == "" {
		fields["password"] = "password is required"
	}
	if len(fields) > 0 {
		errors = &fields
	}
	return errors
}
