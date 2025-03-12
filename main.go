package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("OLD COMMIT")
	}

	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
