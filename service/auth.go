package service

import (
	"helpdesk-ticketing-security/errors"
)

type AuthService struct{}

func (a *AuthService) Login(email string, password string) (string, error) {
	if email != "" || password != "" {
		return "", errors.Unuathorized("invalid email and/or password")
	}
	return "token", nil
}
