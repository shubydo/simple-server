package server_test

// func Test_routes(t *testing.T) {
// 	// Create a new instance of the server
// 	s := &server{
// 		router: http.NewServeMux(),
// 	}

// 	// Create a mock server
// 	srv := httptest.NewServer(s.router)
// 	defer srv.Close()

// 	// Call the routes method
// 	// s.routes()

// 	// Test that the routes are set up correctly
// 	resp, err := http.Get(srv.URL + "/")
// 	if err != nil {
// 		t.Errorf("Error making request: %v", err)
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
// 	}

// }

// func TestRoutes(t *testing.T) {
// 	handler := http.HandlerFunc(routes)

// 	testCases := []struct {
// 		route    string
// 		expected string
// 	}{
// 		{"/", "Welcome to the homepage!"},
// 		{"/about", "About us: We are a company that makes great things!"},
// 		{"/contact", "Contact us: email@example.com"},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.route, func(t *testing.T) {
// 			req, err := http.NewRequest("GET", tc.route, nil)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			rr := httptest.NewRecorder()
// 			handler.ServeHTTP(rr, req)

// 			if status := rr.Code; status != http.StatusOK {
// 				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 			}

// 			if rr.Body.String() != tc.expected {
// 				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tc.expected)
// 			}
// 		})
// 	}
// }

// func TestHandlers(t *testing.T) {
// 	testCases := []struct {
// 		handler  http.Handler
// 		method   string
// 		path     string
// 		expected string
// 	}{
// 		{homeHandler(), "GET", "/", "Welcome to the homepage!"},
// 		{aboutHandler(), "GET", "/about", "About us: We are a company that makes great things!"},
// 		{contactHandler(), "GET", "/contact", "Contact us: email@example.com"},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.path, func(t *testing.T) {
// 			req, err := http.NewRequest(tc.method, tc.path, nil)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			rr := httptest.NewRecorder()
// 			tc.handler.ServeHTTP(rr, req)

// 			if status := rr.Code; status != http.StatusOK {
// 				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 			}

// 			if rr.Body.String() != tc.expected {
// 				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tc.expected)
// 			}
// 		})
// 	}
// }
