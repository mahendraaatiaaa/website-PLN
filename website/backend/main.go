package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World! aku bisa")
	})
	http.ListenAndServe(":8080", nil)

	fmt.Println("Server is running on port 8080")
}
