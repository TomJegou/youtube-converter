package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/TomJegou/youtube-converter-backend/utils"
)

const (
	messageServerError = "Internal server error sorry 😢"
)

func DlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Bad request wants POST got : %s", r.Method), http.StatusBadRequest)
		return
	} else {
		var body struct {
			Url string `json:"url"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, messageServerError, http.StatusInternalServerError)
			log.Printf("json.NewDecoder(r.Body).Decode(&body): %v", err)
			return
		}
		if err := utils.Download(body.Url); err != nil {
			http.Error(w, messageServerError, http.StatusInternalServerError)
			log.Printf("utils.Download(body.Url): %v", err)
			return
		}
		dir, err := os.Open(".")
		if err != nil {
			http.Error(w, messageServerError, http.StatusInternalServerError)
			log.Printf(`os.Open("."): %v`, err)
			return
		}
		defer dir.Close()
		files, err := dir.Readdir(-1)
		if err != nil {
			http.Error(w, messageServerError, http.StatusInternalServerError)
			log.Printf("dir.Readdir(-1): %v", err)
			return
		}
		for _, file := range files {
			if !file.IsDir() && file.Name() != ".gitkeep" && strings.HasSuffix(file.Name(), ".mp3") {
				byteFile, err := os.ReadFile("./" + file.Name())
				if err != nil {
					http.Error(w, messageServerError, http.StatusInternalServerError)
					log.Printf(`os.ReadFile("./" + file.Name()): %v`, err)
					return
				}
				w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", file.Name()))
				w.Header().Set("Content-Type", "audio/mp3")
				w.Header().Set("Content-Length", fmt.Sprintf("%d", file.Size()))
				w.Write(byteFile)
				if err := os.Remove("./" + file.Name()); err != nil {
					http.Error(w, messageServerError, http.StatusInternalServerError)
					log.Printf(`os.Remove("./" + file.Name()): %v`, err)
					return
				}
			}
		}
	}
}
