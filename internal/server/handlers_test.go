package server_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shubydo/simple-server/internal/server"
)

func Test_handleHello(t *testing.T) {
	s := server.New()
	tests := []struct {
		name               string
		method             string
		expectedStatusCode int
		expectedBody       []byte
	}{
		{
			name:               "Return correct response for GET requests to /hello",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedBody:       []byte(`Hello!`),
		},
		{
			name:               "Return error message if other method besides GET is made to /hello",
			method:             http.MethodPost,
			expectedStatusCode: http.StatusMethodNotAllowed,
			expectedBody:       []byte(`Method not allowed`),
		},
	}

	for _, tt := range tests {
		// scopelint: Using a reference for the variable on range scope `tt` (scopelint)
		// https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			// Mock the request
			request, err := http.NewRequest(tc.method, "/hello", nil)
			if err != nil {
				t.Fatalf("Error creating request: %v", err)
			}

			// Mock the handler response
			response := httptest.NewRecorder()
			s.ServeHTTP(response, request)

			actualStatusCode := response.Code
			actualBody, _ := io.ReadAll(response.Body)

			if actualStatusCode != tc.expectedStatusCode {
				t.Fatalf("Actual status code != expected:\nactual: %d\nexpected: %d", actualStatusCode, tc.expectedStatusCode)
			}

			if !bytes.Equal(actualBody, tc.expectedBody) {
				t.Fatalf("Actual response body != expected:\nactual: %s\nexpected: %s", actualBody, tc.expectedBody)
			}
		})
	}
}

func Test_handleIndex(t *testing.T) {
	s := server.New() // Mock server
	tests := []struct {
		name               string
		method             string
		expectedStatusCode int
		expectedBody       []byte
	}{
		{
			name:               "Return correct response for GET requests to /",
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
			expectedBody:       []byte("Index!"),
		},
	}

	for _, tt := range tests {
		tc := tt // scopelint: Using a reference for the variable on range scope `tt` (scopelint)
		t.Run(tc.name, func(t *testing.T) {
			request, err := http.NewRequest(tc.method, "/", nil)
			if err != nil {
				t.Fatalf("Error creating request: %v", err)
			}

			// Mock the handler response
			response := httptest.NewRecorder()
			s.ServeHTTP(response, request)

			actualStatusCode := response.Code
			actualBody, _ := io.ReadAll(response.Body)

			if actualStatusCode != tc.expectedStatusCode {
				t.Fatalf("Actual status code != expected:\nactual: %d\nexpected: %d", actualStatusCode, tc.expectedStatusCode)
			}

			if !bytes.Equal(actualBody, tc.expectedBody) {
				t.Fatalf("Actual response body != expected:\nactual: %s\nexpected: %s", actualBody, tc.expectedBody)
			}
		})
	}
}
