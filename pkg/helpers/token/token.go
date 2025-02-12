package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"net/http"
	"os"
	"strings"
	"time"
)

var jwtSecret = []byte(os.Getenv("AVITO_SECRET"))

type Claims struct {
	Username   string `json:"username"`
	EmployeeID int    `json:"employee_id"`
	jwt.StandardClaims
}

func Generate(username string, employeeID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username:   username,
		EmployeeID: employeeID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Id:        uuid.New().String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func Validate(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}

func ExtractValidToken(r *http.Request) (*Claims, error) {
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		return nil, http.ErrNoCookie
	}
	parts := strings.Split(tokenHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, jwt.ErrSignatureInvalid
	}
	return Validate(parts[1])
}
