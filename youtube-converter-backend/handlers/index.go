package handlers

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my youtube-converter api"))
}
