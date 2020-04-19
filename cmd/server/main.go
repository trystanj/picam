package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Listening on port 8080")

	http.HandleFunc("/image", image)
	http.HandleFunc("/video", video)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// add auth
func image(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "image")
}

func video(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "video")
}
