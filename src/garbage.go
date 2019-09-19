package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {

	// init rgb led
	InitRgbLed()

	// init distance to 3 cm
	isFull := 3

	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	PIN_TRIGGER, err := strconv.Atoi(os.Getenv("PIN_TRIGGER"))
	if err != nil {
		PIN_TRIGGER = 4
	}

	PIN_ECHO, err := strconv.Atoi(os.Getenv("PIN_ECHO"))
	if err != nil {
		PIN_ECHO = 17
	}

	triPin := rpio.Pin(PIN_TRIGGER)
	echoPin := rpio.Pin(PIN_ECHO)

	triPin.Output()
	echoPin.Input()

	// Clean up on ctrl-c and turn lights out
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		triPin.Low()
		echoPin.Low()
		os.Exit(0)
	}()

	defer rpio.Close()

	triPin.Low()
	fmt.Println("Waiting for sensor to settle")
	time.Sleep(1)

	for {

		var startTime, endTime int64

		// fmt.Println("Calculating distance")

		// initialize sensor
		triPin.High()
		time.Sleep(time.Nanosecond)
		triPin.Low()

		// start echoing
		for echoPin.Read() == 0 {
			startTime = time.Now().UnixNano()
		}

		// an echo received
		for echoPin.Read() == 1 {
			endTime = time.Now().UnixNano()
		}

		// calculating result
		// time in nanosec
		// 17000 = (sound speed (320m/s) / 2 converted to cm/s)
		// 1e9 to get result by sec
		distance := (float32(endTime-startTime) * 17000) / 1e9

		if int(distance) > isFull {
			Green()
			fmt.Printf("Distance to be full: %d cm, distance value: %.f \n", (int(distance) - isFull), distance)
		} else {
			Red()
			fmt.Printf("Garbage is full, %.f \n", distance)
		}

		// print result
		time.Sleep(time.Second)

	}

}
