package main

import (
	"net/http"

	"github.com/trystanj/picam/pkg/image"
	"github.com/trystanj/picam/pkg/video"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Listening on port 8080")

	http.HandleFunc("/image", img)
	http.HandleFunc("/video", vid)
	http.HandleFunc("/stream", streamVid)
	http.HandleFunc("/mp4", mp4)
	http.HandleFunc("/capturemp4", captureMp4)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// add auth
func img(w http.ResponseWriter, r *http.Request) {
	output, err := image.Capture(10)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	_, err = w.Write(output)
	if err != nil {
		log.Errorf("Error responding: %w", err)
	}
}

func vid(w http.ResponseWriter, r *http.Request) {
	output, err := video.Capture()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "video/H264")
	_, err = w.Write(output)
	if err != nil {
		log.Errorf("Error responding: %w", err)
	}
}

// this doesn't quite stream, because it streams raw h264 back. browsers evidently can't stream
// raw h264. so need to find a way to encode/stream this on the fly, ideally.
func streamVid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "video/H264")
	video.Stream(w)
}

func mp4(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/home/pi/out.mp4")
}

// we can't immediately serve the file from here, because browsers will use Range requests
// and retrigger the capture/encoding. Instead redirect to the file itself
func captureMp4(w http.ResponseWriter, r *http.Request) {
	file := "out.h264"
	mp4 := "out.mp4"

	video.Save(file)
	video.MP4Box(file, mp4)
	http.Redirect(w, r, "mp4", 303)
}
