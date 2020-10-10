package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/doniacld/notifier/notifier"
	"os"
	"time"
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
	intervalDur, err := time.ParseDuration(*interval)
	if err != nil {
		panic("interval has not a valid format")
	}

	msgsQueue := make(chan string, 5)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		go sender(msgsQueue, message)
		time.Sleep(intervalDur)
		notifier.Notify(*url, msgsQueue)
	}
}

func sender(c chan string, message string) {
	c <- message
}
