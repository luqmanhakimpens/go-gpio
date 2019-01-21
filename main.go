package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	defer rpio.Close()

	led := rpio.Pin(3)
	buzzer := rpio.Pin(2)
	
	buzzer.Output()
	led.Output()

	for {
		led.High()
		buzzer.High()
		time.Sleep(time.Second / 5)
		led.Low()
		buzzer.Low()
		time.Sleep(time.Second * 2)
	}
}
