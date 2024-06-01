package main

import (
	"log"
	"net/http"
)

func fetchAudioURL(url string) string {
	res, err := http.Head(url)

	if err != nil {
		log.Fatal(err)
	}

	contentType := res.Header.Get("Content-Type")

	if contentType == "audio/mpeg" {
		return url
	} else {
		panic("Unsupported media type")
	}
}
