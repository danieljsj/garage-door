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

var latestKnownTriggeringTime uint64

func main() {
	myinit()

	// theoretically this should go forever, but for now we're not going to. don't want to make FB mad by accident
	for i := 0; i < 10; i++ {
		spoofLatestTriggering()

		latestTriggering := getLatestTriggering()
		if latestTriggering.Time > latestKnownTriggeringTime {
			latestTriggering.Time = latestKnownTriggeringTime
			fire()
		}
		time.Sleep(time.Second * 1)
	}


}

func myinit() {
	prepPins()
	latestTriggering := getLatestTriggering()
	latestKnownTriggeringTime = latestTriggering.Time
}


// GPIO CONTROL:

// Use mcu pin 17, corresponds to physical pin 11 on the pi
var pin = rpio.Pin(17)

func prepPins() {
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

}
func fire() {
	pin.Low()
	time.Sleep(time.Second / 5)
	pin.High()
}



// WEBSERVER:

func serve() {
	http.HandleFunc("/trigger", triggerHandler)
	http.ListenAndServe(":8080", nil)
}

func triggerHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path[1:]
	fire()
}



// FIREBASE:

type Triggering struct {
	Time uint64
	Username string
}


func getLatestTriggering() Triggering {
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

    latestTriggeringRef := firebase.NewReference(latestTriggeringUrl).Auth(authToken).Export(false)

    latestTriggering := Triggering{}

    if err = latestTriggeringRef.Value(latestTriggering); err != nil {
        panic(err)
    }

    fmt.Println("latestTriggering.Time")
    fmt.Println(latestTriggering.Time)
    fmt.Println("latestTriggering.Username")
    fmt.Println(latestTriggering.Username)

    return latestTriggering
}

func spoofLatestTriggering() bool {
    var err error

    latestTriggeringUrl := "https://garage-opener.firebaseIO.com/latestTriggering"

    // Can also be your Firebase secret: (CURRENTLY IS. THAT'S OKAY, BECAUSE THIS IS SITTING ON MY GARAGE MACHINE, AND EVENTUALLY ONLY IN BINARY, SO NOT WORRIED)
    authToken := "TDNkOlwWlXMZqFaGntBVrRE819MbPrcZdsFRaO3K"

    latestTriggeringRef := firebase.NewReference(latestTriggeringUrl).Auth(authToken).Export(false)

    // Create the value.
    personName := Triggering{
        Time: latestKnownTriggeringTime + uint64(1),
        Username:  "rebekahsj",
    }

    // Write the value to Firebase.
    if err = latestTriggeringRef.Write(personName); err != nil {
        panic(err)
    }
    return false
}