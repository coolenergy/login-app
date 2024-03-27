package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Cerebrovinny/login-app/config"
	"github.com/Cerebrovinny/login-app/models"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func createTestUser() (models.User, error) {
	db, err := config.GetDatabase()
	if err != nil {
		return models.User{}, err
	}

	client, err := config.GetMongoClient()
	if err != nil {
		return models.User{}, err
	}
	defer client.Disconnect(context.Background())

	collection := db.Collection("users")
	testUser := models.User{
		Username: "testuser",
		Password: models.HashPassword("testpassword"),
	}
	_, err = collection.InsertOne(context.Background(), testUser)
	return testUser, err
}

func TestLoginHandler(t *testing.T) {
	testUser, err := createTestUser()
	if err != nil {
		t.Fatalf("Error creating test user: %v", err)
	}

	tests := []struct {
		name           string
		method         string
		credentials    models.Credentials
		expectedStatus int
		checkJWT       bool
		jwtKey         []byte
	}{
		{
			name:           "Invalid method",
			method:         "GET",
			credentials:    models.Credentials{},
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:   "Invalid payload",
			method: "POST",
			credentials: models.Credentials{
				Username: "nonexistent",
				Password: "testpassword",
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:   "Wrong password",
			method: "POST",
			credentials: models.Credentials{
				Username: testUser.Username,
				Password: "wrongpassword",
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:   "Valid login with JWT check",
			method: "POST",
			credentials: models.Credentials{
				Username: testUser.Username,
				Password: "testpassword",
			},
			expectedStatus: http.StatusOK,
			checkJWT:       true,
		},
		{
			name:   "Valid login with wrong JWT",
			method: "POST",
			credentials: models.Credentials{
				Username: testUser.Username,
				Password: "testpassword",
			},
			expectedStatus: http.StatusOK,
			jwtKey:         []byte("wrong_jwt_key"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			payload, _ := json.Marshal(test.credentials)
			req, err := http.NewRequest(test.method, "/login", bytes.NewBuffer(payload))
			if err != nil {
				t.Fatalf("Error creating request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(LoginHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != test.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v, expected %v", status, test.expectedStatus)
			}

			if test.checkJWT {
				tokenString := rr.Header().Get("Set-Cookie")
				tokenString = strings.TrimPrefix(tokenString, "token=")
				tokenString = strings.Split(tokenString, ";")[0]

				claims := &Claims{}
				token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
					return jwtKey, nil
				})

				if err != nil {
					t.Errorf("Error parsing JWT token: %v", err)
				}

				if !token.Valid {
					t.Errorf("JWT token is invalid")
				}

				if claims.Username != test.credentials.Username {
					t.Errorf("JWT claims contain wrong username: got %v, expected %v", claims.Username, test.credentials.Username)
				}
			}
		})
	}
}
