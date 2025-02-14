package handlers

import (
	"encoding/json"
	"github.com/mrkucher83/avito-shop/internal/models"
	"net/http"
)

func (rp *Repo) GetEmployeeInfo(w http.ResponseWriter, r *http.Request) {
	employeeID, err := GetEmployeeIDFromContext(r)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	var resp models.EmployeeInfoResponse

	// Get coins of employee
	resp.Coins, err = rp.storage.GetCoinsById(r.Context(), employeeID)
	if err != nil {
		http.Error(w, "Error getting coin balance: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Get employee's inventory
	resp.Inventory.Items, err = rp.storage.GetInventoryById(r.Context(), employeeID)
	if err != nil {
		http.Error(w, "Error getting employee's purchases: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Get employee's coin history
	resp.CoinsHistory.Received.Items, err = rp.storage.GetReceivedCoins(r.Context(), employeeID)
	if err != nil {
		http.Error(w, "Error getting employee's received coins: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp.CoinsHistory.Sent.Items, err = rp.storage.GetSentCoins(r.Context(), employeeID)
	if err != nil {
		http.Error(w, "Error getting employee's sent coins: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
