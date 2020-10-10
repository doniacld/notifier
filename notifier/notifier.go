/* Package notifier holds
 */
package notifier

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

// Notify notifies that a message arrives by sending this message to the given URL
// TODO DONIA should manage a lot of requests at the same time
func Notify(url string, msgsQueue chan string) error {
	// send a request to the given URL
	for i := 0; i < 5; i++ {
		go worker(url, msgsQueue)
	}
	return nil
}

func worker(url string, msgsQueue chan string) {
	msg := <-msgsQueue
	_ = httptest.NewRequest(http.MethodPost, url, strings.NewReader(msg))
	fmt.Println(msg)
}
