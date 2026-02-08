package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TomJegou/youtube-converter-backend/handlers"
)

const (
	port = "46460"
)

func main() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/dl", handlers.DlHandler)
	fmt.Printf("API is up ! 🔥\nOn port 0.0.0.0:%s: http://localhost:%s\n", port, port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Printf(`http.ListenAndServe(fmt.Sprintf(":%%s", port), nil): %v`, err)
	}
}
