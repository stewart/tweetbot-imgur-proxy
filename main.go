package main

import (
	"log"
	"net/http"
	"os"
)

const ENDPOINT = "https://api.imgur.com/3/image"

var (
	CLIENT_ID = os.Getenv("IMGUR_CLIENT_ID")
	PORT      = os.Getenv("PORT")
)

func init() {
	if PORT == "" {
		PORT = "3000"
	}

	if CLIENT_ID == "" {
		log.Fatal("IMGUR_CLIENT_ID not set")
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Print("server starting")
	http.ListenAndServe(":"+PORT, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)

	if r.Method != "POST" {
		log.Println("Invalid (non-POST) request, HTTP 405")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	json, err := uploadAttachedFile(r)

	if err != nil {
		log.Println("  error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
