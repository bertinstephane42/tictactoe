package main

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// Lecture d'un fichier sonore
func playSound(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier audio :", err)
		return
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		fmt.Println("Erreur lors du décodage audio :", err)
		return
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
}
