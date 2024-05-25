package main

// server

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! %s", time.Now())
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server is running on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server failed to start", err)
	}
}
