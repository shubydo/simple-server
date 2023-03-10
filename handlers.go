package main

// Go did not recognize/resolve  initHandlers() not being able to be found in the main package
// func InitHandlers() {
// 	http.HandleFunc("/hello", handleHello)
// }

// func handleHello(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		fmt.Fprintf(w, "Hello!")
// 		return
// 	}

// 	// Return "Method not allowed" if not a GET request
// 	fmt.Fprintf(w, "Method not allowed")
// }
