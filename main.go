package main

import (
	"log"
	"net/http"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
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

func main() {
	audioURL := fetchAudioURL("https://stream1.cprnetwork.org/cpr1_lo")

	resp, err := http.Get(audioURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	streamer, format, err := mp3.Decode(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
