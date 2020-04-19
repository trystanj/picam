package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/trystanj/picam/pkg/image"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Listening on port 8080")

	http.HandleFunc("/image", img)
	http.HandleFunc("/video", video)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// add auth
func img(w http.ResponseWriter, r *http.Request) {
	output, err := image.Image("arg")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result := bytes.NewBuffer(output).String()
	fmt.Fprintf(w, result)
}

func video(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "video")
}
