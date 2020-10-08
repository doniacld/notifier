package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"
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
	_, err := time.ParseDuration(*interval)
	if err != nil {
		panic("interval has not a valid format")
	}
	wg := sync.WaitGroup{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		wg.Add(1)
		go notifier.Notify(*url, message)
	}
	wg.Wait()
}
