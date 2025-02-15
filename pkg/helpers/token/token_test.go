package token

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	username := "testUser"
	employeeID := 12345

	token, err := Generate(username, employeeID)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	claims, err := Validate(token)
	require.NoError(t, err)
	require.NotNil(t, claims)
	require.Equal(t, username, claims.Username)
	require.Equal(t, employeeID, claims.EmployeeID)
}

func TestValidateInvalidToken(t *testing.T) {
	invalidToken := "invalid.token.string"

	claims, err := Validate(invalidToken)
	require.Error(t, err)
	require.Nil(t, claims)
}

func TestExtractValidToken(t *testing.T) {
	token, _ := Generate("testUser", 12345)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	claims, err := ExtractValidToken(req)
	require.NoError(t, err)
	require.NotNil(t, claims)
	require.Equal(t, "testUser", claims.Username)
}

func TestExtractInvalidToken(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer invalid.token.string")

	claims, err := ExtractValidToken(req)
	require.Error(t, err)
	require.Nil(t, claims)
}
