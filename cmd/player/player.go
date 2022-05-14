package player

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func PlaySound() {
	filePath := os.Getenv("START_VOICE_FILE_PATH")
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	streamer, format, err := wav.Decode(f)
	if err != nil {
		panic(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(
		streamer,
		beep.Callback(
			func() {
				done <- true
			},
		),
	))
	<-done
}
