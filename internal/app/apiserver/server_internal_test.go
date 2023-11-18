package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"rest-api/internal/app/model"
	"rest-api/internal/app/nlab/testnlab"
	"testing"
)

func TestServer_AuthenticateUser(t *testing.T) {
	nlab := testnlab.New()
	u := model.TestUser(t)
	nlab.User().Create(u)
	testCases := []struct {
		name         string
		cookieValue  map[interface{}]interface{}
		expectedCode int
	}{
		{
			name: "authenticated",
			cookieValue: map[interface{}]interface{}{
				"user_id": u.ID,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "not authenticated",
			cookieValue:  nil,
			expectedCode: http.StatusUnauthorized,
		},
	}

	secretKey := []byte("secret")
	s := newServer(nlab, sessions.NewCookieStore(secretKey))
	sc := securecookie.New(secretKey, nil)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			cookieStr, _ := sc.Encode(sessionName, tc.cookieValue)
			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))
			s.authenticateUser(handler).ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)

		})
	}
}

func TestServer_handleUsersCreate(t *testing.T) {
	s := newServer(testnlab.New(), sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"login":    "Test1",
				"password": "password",
				"name":     "Test1",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid paylot",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"login":    "Test2",
				"password": "",
				"name":     "Test2",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

	//rec := httptest.NewRecorder()
	//req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	//s := newServer(testnlab.New())
	//s.ServeHTTP(rec, req)
	//assert.Equal(t, rec.Code, http.StatusOK)

}
func TestServer_HandleSessionsCreate(t *testing.T) {
	u := model.TestUser(t)
	nlab := testnlab.New()
	nlab.User().Create(u)
	s := newServer(nlab, sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"login":    u.Login,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "valid",
			payload: map[string]string{
				"login":    "admin",
				"password": "admin",
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid paylod",
			payload:      "payload",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid login",
			payload: map[string]string{
				"login":    "kosssstia",
				"password": "admin",
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid pass",
			payload: map[string]string{
				"login":    u.Login,
				"password": "admin",
			},
			expectedCode: http.StatusUnauthorized,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
