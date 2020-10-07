package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/doniacld/notifier/notifier"
)

const (
	helpFlag          = "help"
	intervalShortFlag = "i"
	intervalFlag      = "interval"
	intervalDefault   = "5s"
	urlFlag           = "url"
	urlDefault        = "http://localhost:8080/notify"
)

func main() {
	fmt.Println("Let's start to notify")

	// flags definition
	flag.String(helpFlag, "", "Show context-sensitive help (also\ntry --help-long and --help-man).")
	url := flag.String(urlFlag, urlDefault, "Notification URL")
	if url == nil {
		panic("url is empty")
	}
	interval := flag.String(intervalShortFlag, intervalDefault, "Notification interval")
	fmt.Println(*interval)
	flag.Parse()

	// convert interval into a int
	intervalInt := 5

	// TODO DONIA
	// retrieve the message
	message := "coucou toi!"

	// enter the loop
	sender(*url, intervalInt, message)
}

// sender sends the provided message at the given internal
func sender(url string, interval int, message string) {
	for {
		time.Sleep(time.Duration(interval) * time.Second)
		go notifier.Notify(url, message)
	}
}
