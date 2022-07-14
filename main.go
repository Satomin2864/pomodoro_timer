package main

import (
	"strconv"

	"pomodoro/cmd/console"
	"pomodoro/cmd/player"
	"pomodoro/cmd/timer"
)

func main() {
	workTime, _ := console.Console()
	// TODO: ctrl+c で終了しても、後続処理が動いているので止める必要あり
	intWorkTime, _ := strconv.Atoi(workTime)
	timer.Timer(intWorkTime)
	player.PlaySound()
}
