package controller

import (
	"encoding/json"
	"fmt"
	"helpdesk-ticketing-security/errors"
	"helpdesk-ticketing-security/model"
	"helpdesk-ticketing-security/service"
	"net/http"
)

type AuthController struct {
	Service service.AuthService
}

func (authController *AuthController) Login(res http.ResponseWriter, req *http.Request) {
	var body model.AuthLoginRequest

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		http.Error(res, "failed to read request body", http.StatusBadRequest)
		return
	}
	if err := body.Validate(); err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusUnprocessableEntity)
		res.Write([]byte(err.Error()))
		return
	}
	if token, err := authController.Service.Login(body.Email, body.Password); err == nil {
		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res, token)

	} else {
		var status int

		switch err.(type) {
		case errors.Unuathorized:
			status = http.StatusUnauthorized
		default:
			status = http.StatusInternalServerError
		}

		http.Error(res, err.Error(), status)
	}
}

func (a *AuthController) User(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(res, "NOT IMPLEMENTED")
}

func (a *AuthController) Refresh(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(res, "NOT IMPLEMENTED")
}
