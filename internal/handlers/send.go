package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/mrkucher83/avito-shop/internal/godb"
	"github.com/mrkucher83/avito-shop/internal/models"
	"net/http"
)

func (rp *Repo) SendCoin(w http.ResponseWriter, r *http.Request) {
	employeeID, err := GetEmployeeIDFromContext(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Extract data from request body
	var req models.SendCoinRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.ToUser == "" || req.Amount == 0 {
		http.Error(w, "Receiver's name and amount required", http.StatusBadRequest)
		return
	}

	// Check if there are enough funds
	coins, err := rp.storage.GetCoinsById(r.Context(), employeeID)
	if err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	if coins < req.Amount {
		http.Error(w, "Insufficient funds", http.StatusPaymentRequired)
		return
	}

	// Get a receiver
	receiver, err := rp.storage.GetEmployee(r.Context(), req.ToUser)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "There is no employee with that name", http.StatusBadRequest)
			return
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
	}

	// Transaction starts
	tx, err := rp.storage.BeginTransaction(r.Context())
	if err != nil {
		http.Error(w, "Transaction error", http.StatusInternalServerError)
		return
	}

	defer func() {
		if err != nil {
			godb.RollbackOnError(r.Context(), tx)
			http.Error(w, "Transaction failed, no coins were transferred", http.StatusInternalServerError)
		}
	}()

	// Deduct coins from sender
	if err = rp.storage.ReduceBalance(r.Context(), tx, req.Amount, employeeID); err != nil {
		http.Error(w, "Failed to deduct coins", http.StatusInternalServerError)
		return
	}

	// Add coins to receiver
	if err = rp.storage.IncreaseBalance(r.Context(), tx, req.Amount, receiver.ID); err != nil {
		http.Error(w, "Failed to add coins", http.StatusInternalServerError)
		return
	}

	// Insert transaction record
	if err = rp.storage.InsertTransaction(r.Context(), tx, employeeID, receiver.ID, req.Amount); err != nil {
		http.Error(w, "Failed to record transaction", http.StatusInternalServerError)
		return
	}

	// Transaction completion
	err = tx.Commit(r.Context())
	if err != nil {
		http.Error(w, "Transaction commit failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(models.Response{Message: "Transaction successful"}); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
