package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
	flag.Parse()

	// convert interval into a duration format
	intervalDuration, err := time.ParseDuration(*interval)
	if err != nil {
		panic("interval has not a valid format")
	}

	// Parse stdin and sends the messages
	sender(*url, intervalDuration)
}

// sender sends the provided message at the given internal
func sender(url string, interval time.Duration) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			time.Sleep(interval)
			message := scanner.Text()
			go notifier.Notify(url, message)
		}
	}
}
