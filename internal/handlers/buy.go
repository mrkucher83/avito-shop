package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/mrkucher83/avito-shop/internal/models"
	"github.com/mrkucher83/avito-shop/pkg/helpers/token"
	"net/http"
)

func (rp *Repo) BuyItem(w http.ResponseWriter, r *http.Request) {
	employeeID, err := GetEmployeeIDFromContext(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	itemName := chi.URLParam(r, "item")
	if itemName == "" {
		http.Error(w, "Item name is required", http.StatusBadRequest)
		return
	}

	item, err := rp.storage.GetMerchByName(r.Context(), itemName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Item not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	coins, err := rp.storage.GetCoinsById(r.Context(), employeeID)
	if err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	if coins < item.Price {
		http.Error(w, "Insufficient funds", http.StatusPaymentRequired)
		return
	}

	// Transaction starts
	tx, err := rp.storage.BeginTransaction(r.Context())
	if err != nil {
		http.Error(w, "Transaction error", http.StatusInternalServerError)
		return
	}

	if err = rp.storage.UpdateBalance(r.Context(), tx, item.Price, employeeID); err != nil {
		http.Error(w, "Failed to update balance", http.StatusInternalServerError)
		return
	}

	if err = rp.storage.RecordPurchase(r.Context(), tx, employeeID, item.ID, 1); err != nil {
		http.Error(w, "Failed to record purchase", http.StatusInternalServerError)
		return
	}

	// Transaction completion
	err = tx.Commit(r.Context())
	if err != nil {
		tx.Rollback(r.Context()) // Откатываем, если commit не удался
		http.Error(w, "Transaction commit failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(models.Response{Message: "Purchase successful"}); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetEmployeeIDFromContext(r *http.Request) (int, error) {
	claims, ok := r.Context().Value("userClaims").(*token.Claims)
	if !ok || claims == nil {
		return 0, errors.New("employee ID not found in context")
	}
	return claims.EmployeeID, nil
}
