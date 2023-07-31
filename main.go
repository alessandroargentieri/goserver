package main

import (
  "net/http"
  "log"
)

func main() {
   http.HandleFunc("/hello",  func(w http.ResponseWriter, r *http.Request){
   
	// Add CORS headers to allow requests from any origin
     	w.Header().Set("Access-Control-Allow-Origin", "*")
     	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
     	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

	// Check if the request is an OPTIONS preflight request
     	if r.Method == "OPTIONS" {
		// Respond with 200 OK for preflight requests
		return
     	}

      	w.Write([]byte("Hello world!"))
   })
   log.Fatal(http.ListenAndServe(":8080", nil))
   log.Println("HTTP Server inited on port 8080")
}
