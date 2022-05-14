package timer

import (
	"fmt"
	"time"
)

const (
	workTimeMinutes  = 4 * time.Second
	breakTimeMinutes = 5
)

func Timer() {
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
