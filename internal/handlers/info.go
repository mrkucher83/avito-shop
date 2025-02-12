package handlers

import (
	"github.com/mrkucher83/avito-shop/pkg/logger"
	"net/http"
)

func (rp *Repo) GetEmployeeInfo(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Welcome to Info Page!")); err != nil {
		logger.Warn("failed to write response: ", err)
	}
}
