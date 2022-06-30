package controller

import (
	"encoding/json"
	"fmt"
	"helpdesk-ticketing-security/errors"
	"helpdesk-ticketing-security/model"
	"helpdesk-ticketing-security/service"
	"net/http"
)

type PasswordController struct {
	Service service.PasswordService
}

func (passwordController *PasswordController) Forgot(res http.ResponseWriter, req *http.Request) {
	var body model.ForgotPasswordRequest

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if err := body.Validate(); err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusUnprocessableEntity)
		res.Write([]byte(err.Error()))
		return
	}
	if err := passwordController.Service.SendResetCode(body.Email); err == nil {
		res.WriteHeader(http.StatusOK)
		res.Write(nil)

	} else {
		var status int

		switch err.(type) {
		case errors.BadRequest:
			status = http.StatusBadRequest
		default:
			status = http.StatusInternalServerError
		}

		http.Error(res, err.Error(), status)
	}
}

func (p *PasswordController) Reset(res http.ResponseWriter, req *http.Request) {
	var body model.ResetPasswordRequest

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if err := body.Validate(); err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusUnprocessableEntity)
		res.Write([]byte(err.Error()))
		return
	}

	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, "NOT IMPLEMENTED")
}
