package main

import (
	"image/color"
	"pomodoro/cmd/player"
	"pomodoro/cmd/timer"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// タイマーを止めれるようにする必要あり？

func main() {
	timer.Timer()
	player.PlaySound()
	appCreate()
}
func appCreate() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text2.Move(fyne.NewPos(20, 20))
	content := container.NewWithoutLayout(text1, text2)
	// content := container.New(layout.NewGridLayout(2), text1, text2)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

// func updateTime(clock *widget.Label) {
// 	formatted := time.Now().Format("Time: 03:04:05")
// 	clock.SetText(formatted)
// }
