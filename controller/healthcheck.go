package controller

import (
	"fmt"
	"net/http"
)

func Ping(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "SERVICE RUNNING")
}
