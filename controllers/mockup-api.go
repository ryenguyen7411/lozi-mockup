package controllers

import (
	"fmt"
	"net/http"
)

// MockupAPIHandler ...
func MockupAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/")

	fmt.Println("MUX", r.URL, r.Method)
	w.Write([]byte("foo"))
}

// func mockupAPIPost() {

// }

// func mockupAPIPut() {

// }

// func mockupAPIDelete() {

// }
