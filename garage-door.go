package main

import (
	"fmt"
	"os"
	"time"
	// "log"
	
	"net/http"

	"github.com/stianeikeland/go-rpio"
    "github.com/melvinmt/firebase"
)

var (
	// Use mcu pin 17, corresponds to physical pin 11 on the pi
	pin = rpio.Pin(17)
)

func mainDISABLED() {
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

type Triggering struct {
	Time uint64
	Username string
}


func main() {
    var err error

    latestTriggeringUrl := "https://garage-opener.firebaseIO.com/latestTriggering"

    // Can also be your Firebase secret: (CURRENTLY IS. THAT'S OKAY, BECAUSE THIS IS SITTING ON MY GARAGE MACHINE, AND EVENTUALLY ONLY IN BINARY, SO NOT WORRIED)
    authToken := "TDNkOlwWlXMZqFaGntBVrRE819MbPrcZdsFRaO3K"

    // Auth is optional:
    // ref := firebase.NewReference(latestTriggeringUrl).Auth(authToken)

    // Create the value.
    // personName := PersonName{
    //     First: "Fred",
    //     Last:  "Swanson",
    // }

    // Write the value to Firebase.
    // if err = ref.Write(personName); err != nil {
    //     panic(err)
    // }

    // Now, we're going to retrieve the person.
    // personlatestTriggeringUrl := "https://SampleChat.firebaseIO.com/users/fred"

    // personRef := firebase.NewReference(personlatestTriggeringUrl).Export(false)
    latestTriggeringRef := firebase.NewReference(latestTriggeringUrl).Auth(authToken).Export(false)

    var latestTriggering *Triggering = &Triggering{}

    if err = latestTriggeringRef.Value(latestTriggering); err != nil {
        panic(err)
    }

    fmt.Println("let's see what we got:")
    fmt.Println(latestTriggering.Time)
    fmt.Println(latestTriggering.Username)
}