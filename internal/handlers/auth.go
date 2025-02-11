package handlers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/mrkucher83/avito-shop/internal/models"
	"github.com/mrkucher83/avito-shop/pkg/helpers/hasher"
	"net/http"
	"time"
)

var jwtSecret = []byte("avito_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (rp *Repo) SignUp(w http.ResponseWriter, r *http.Request) {
	var req models.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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
		if err.Error() == "no rows in result set" {
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
		if err := rp.storage.CreateEmployee(r.Context(), req); err != nil {
			http.Error(w, "Error creating employee", http.StatusInternalServerError)
			return
		}
	} else {
		// Check password for existing employee
		if err := hasher.CheckPasswordHash(req.Password, existingEmployee.Password); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
	}

	// Generate JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: req.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Id:        uuid.New().String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(models.AuthResponse{Token: tokenString}); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
