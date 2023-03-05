package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Set the response content type to HTML
			w.Header().Set("Content-Type", "text/html")
			// Write the HTML response to the client
			fmt.Fprint(w, "<html><body><h1>Hello, World!</h1></body></html>")
		case http.MethodPost:
			// Set the response content type to JSON
			w.Header().Set("Content-Type", "application/json")
			// Parse the request body as JSON
			var reqBody struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}
			err := json.NewDecoder(r.Body).Decode(&reqBody)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// Create the response JSON
			resp := map[string]interface{}{
				"message": fmt.Sprintf("Hello, %s! You are %d years old.", reqBody.Name, reqBody.Age),
			}
			// Write the JSON response to the client
			err = json.NewEncoder(w).Encode(resp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		default:
			// Set the response content type to plain text
			w.Header().Set("Content-Type", "text/plain")
			// Write a default response to the client
			fmt.Fprint(w, "Unsupported HTTP method")
		}
	})

	// Start the server and listen on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
