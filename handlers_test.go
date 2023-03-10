package main

// func TestHandleHello(t *testing.T) {
// 	tests := []struct {
// 		name          string
// 		requestMethod string
// 		expected      string
// 	}{
// 		{
// 			name:          "Return correct response for GET requests to /hello",
// 			requestMethod: http.MethodGet,
// 			expected:      "Hello!",
// 		},
// 		{
// 			name:          "Return error message if other method besides GET is made to /hello",
// 			requestMethod: http.MethodPost,
// 			expected:      "Method not allowed",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			request, _ := http.NewRequest(tt.requestMethod, "/hello", nil)
// 			response := httptest.NewRecorder()

// 			handleHello(response, request)

// 			actual := response.Body.String()
// 			if actual != tt.expected {
// 				t.Fatalf("Actual response body != expected:\nactual:%q\nexpected: %q", actual, tt.expected)
// 			}
// 		})
// 	}
// }
