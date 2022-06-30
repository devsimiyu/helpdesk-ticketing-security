package service

import (
	"fmt"
	"helpdesk-ticketing-security/errors"
)

type PasswordService struct{}

func (p *PasswordService) SendResetCode(email string) error {
	if email != "" {
		return errors.BadRequest(
			fmt.Sprintf(
				"No account is associated with this email: %s",
				email,
			),
		)
	}
	return nil
}
