package main

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

const (
	workTimeMinutes  = 4 * time.Second
	breakTimeMinutes = 5
)

// タイマーを止めれるようにする必要あり？

func main() {
	timer()
	playSaund()
}
func playSaund() {
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
func timer() {
	fmt.Println("timer start")
	elapsedTime := 0
	nowTime := workTimeMinutes - (time.Duration(elapsedTime) * time.Second)
	fmt.Println(nowTime)
	timer1 := time.NewTimer(workTimeMinutes)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			elapsedTime++
			nowTime = workTimeMinutes - (time.Duration(elapsedTime) * time.Second)
			fmt.Println(nowTime)
		}
	}()
	<-timer1.C
	nowTime = time.Duration(0)
	fmt.Println(nowTime)
	fmt.Println("timer end")
}
