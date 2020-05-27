package controllers

import (
	"fmt"
	"net/http"
)

// MockupAPIHandler ...
/** TODO: phase 2: consistent return
- stored fake data into db for next use
*/
func MockupAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: URL, Method
	/**
	- Step 1: get mockup config from db. If failed -> Fatal
	- Step 2: parse data model to JSON
	- Step 3: fake data from data model
	- Step 4: return
	*/

	fmt.Println("MUX", r.URL, r.Method)
	w.Write([]byte("foo"))
}

// func mockupAPIPost() {

// }

// func mockupAPIPut() {

// }

// func mockupAPIDelete() {

// }
