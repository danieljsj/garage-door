package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
	"net/http"
)

var (
	// Use mcu pin 17, corresponds to physical pin 11 on the pi
	pin = rpio.Pin(17)
)

func main() {
	// Unmap gpio memory when done
	defer rpio.Close()

	// Open and map memory to access gpio, check for errors
	err := rpio.Open(); 
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Set pin to output mode as opposed to input
	pin.Output()

	http.HandleFunc("/trigger", triggerHandler)
	http.ListenAndServe(":8080", nil)
}

func triggerHandler(w http.ResponseWriter, r *http.Request) {

	// r.URL.Path[1:]
	pin.Low()
	time.Sleep(time.Second / 5)
	pin.High()

}