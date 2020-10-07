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
func Notify(url string, message string) error {
	// send a request to the given URL
	_ = httptest.NewRequest(http.MethodPost, url, strings.NewReader(message))
	fmt.Println(message)
	return nil
}
