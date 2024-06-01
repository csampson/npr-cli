package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

func Play(url string, done chan bool) {
	audioURL := fetchAudioURL(url)

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
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func Stop() {
	speaker.Clear()
}
