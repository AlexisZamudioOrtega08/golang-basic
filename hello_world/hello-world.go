package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func main() {
	// Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	// starting server
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Home page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["msg"] = "Status Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error while encoding json. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}
