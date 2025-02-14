package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/mrkucher83/avito-shop/internal/models"
	"github.com/mrkucher83/avito-shop/pkg/helpers/hasher"
	"github.com/mrkucher83/avito-shop/pkg/helpers/token"
	"net/http"
	"time"
)

func (rp *Repo) SignUp(w http.ResponseWriter, r *http.Request) {
	//Check if there is JWT from request Header
	claims, err := token.ExtractValidToken(r)
	if err == nil && claims.ExpiresAt > time.Now().Unix() {
		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(models.AuthResponse{Token: r.Header.Get("Authorization")[7:]}); err != nil {
			http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
			return
		}
		return
	} else if err != nil && err != http.ErrNoCookie && err != jwt.ErrSignatureInvalid {
		http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
		return
	}

	// Extract data from request body
	var req models.AuthRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password required", http.StatusBadRequest)
		return
	}

	// Check if employee exists
	existingEmployee, err := rp.storage.GetEmployee(r.Context(), req.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			existingEmployee = nil
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
	}

	var hashedPassword string
	if existingEmployee == nil {
		// New employee, hash password
		hashedPassword, err = hasher.HashPassword(req.Password)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}
		req.Password = hashedPassword

		// Create employee in database
		if err = rp.storage.CreateEmployee(r.Context(), req); err != nil {
			http.Error(w, "Error creating employee", http.StatusInternalServerError)
			return
		}
	} else {
		// Check password for existing employee
		if err = hasher.CheckPasswordHash(req.Password, existingEmployee.Password); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
	}

	// Generate JWT token
	tokenString, err := token.Generate(req.Username, existingEmployee.ID)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(models.AuthResponse{Token: tokenString}); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
