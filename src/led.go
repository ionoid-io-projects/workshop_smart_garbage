/*

A blinker example using go-rpio library.
Requires administrator rights to run

Toggles a LED on physical pin 19 (mcu pin 10)
Connect a LED with resistor from pin 19 to ground.

*/

package main

import (
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio"
)

var (
	// Use mcu pin 10, corresponds to physical pin 19 on the pi
	pinRed   = rpio.Pin(25)
	pinGreen = rpio.Pin(24)
	pinBlue  = rpio.Pin(23)
)

func InitRgbLed() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	pinRed.Output()
	pinGreen.Output()
	pinBlue.Output()
}

func Red() {
	pinRed.High()
	pinGreen.Low()
	pinBlue.Low()
}

func Green() {
	pinRed.Low()
	pinGreen.High()
	pinBlue.Low()
}

func Orange() {
	pinRed.High()
	pinGreen.Low()
	pinBlue.High()
}

func Init() {
	pinRed.Low()
	pinGreen.Low()
	pinBlue.Low()
}
