package model

import (
	"helpdesk-ticketing-security/errors"
	"net/mail"
)

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

func (f *ForgotPasswordRequest) Validate() error {
	fields, errors := make(errors.ValidationError), error(nil)

	if f.Email == "" {
		fields["email"] = "email is required"

	} else if _, err := mail.ParseAddress(f.Email); err != nil {
		fields["email"] = "email must be valid"
	}
	if len(fields) > 0 {
		errors = &fields
	}
	return errors
}

type ResetPasswordRequest struct {
	Code              string `json:"code"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmedPassword"`
}

func (r *ResetPasswordRequest) Validate() error {
	fields, errors := make(errors.ValidationError), error(nil)

	if r.Code == "" {
		fields["code"] = "code is required"
	}
	if r.Email == "" {
		fields["email"] = "email is required"

	} else if _, err := mail.ParseAddress(r.Email); err != nil {
		fields["email"] = "email must be valid"
	}
	if r.Password == "" {
		fields["password"] = "password is required"

	} else if r.Password != r.ConfirmedPassword {
		fields["password"] = "password do not match"
	}
	if len(fields) > 0 {
		errors = &fields
	}
	return errors
}
