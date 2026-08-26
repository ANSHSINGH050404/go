package main

import (
	"fmt"
	"net/http"
)



func main() {

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Health check OK!")
		w.Write([]byte("ok"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}