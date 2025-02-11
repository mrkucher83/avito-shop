package employee

import (
	"github.com/mrkucher83/avito-shop/pkg/logger"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Sign-Up page of Avito Shop!")); err != nil {
		logger.Warn("failed to write response: ", err)
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Sign-In page of Avito Shop!")); err != nil {
		logger.Warn("failed to write response: ", err)
	}
}
