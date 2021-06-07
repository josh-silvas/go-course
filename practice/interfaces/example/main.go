// Example usage of interfaces to abstract the behavior of an Alert based on the
// type used in the caller.
package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/josh5276/go-course/common/settings"
)

// Alerter defines the behavior of an interface. The Alerter interface will implement a method called
// Alert that does not take an argument and returns an error. Any type that implements this type
// of behavior will satisfy the Alerter interface.
type Alerter interface {
	Alert() error
}

func Send(a Alerter) error {
	return a.Alert()
}

type Slack struct {
	string
}

func (s *Slack) Alert() error {
	//
	// Do some logic to initiate an API call to Slack and send the message to
	// the specified channel.
	//

	url := fmt.Sprintf("https://hooks.slack.com/services/%s", settings.SlackKey)
	// Make a io.Reader type using a CTK query payload
	payload := strings.NewReader(fmt.Sprintf(`{"text": "%s"}`, s.string))

	// NewRequest returns a new Request given a method, URL, and optional body.
	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return err
	}

	if _, err := http.DefaultClient.Do(req); err != nil {
		return err
	}

	fmt.Printf("Sent Slack alert successfully.\nMessage: %s\n\n", s.string)
	return nil
}

type NewRelic struct {
	string
}

func (n *NewRelic) Alert() error {
	//
	// Do some logic to initiate an API call to NewRelic and send the message to
	// the specified account number.
	//

	nr, err := newRelic.NewClient(settings.NRAPIKey, settings.NRAccount)
	if err != nil {
		return err
	}

	r, err := nr.Insights.JournalEvent("GOC", "Examples", map[string]string{"interface_test": n.string})
	if err != nil {
		return err
	}

	// Results stored at:
	// https://insights.newrelic.com/accounts/929627/dashboards/902079
	fmt.Printf("Sent NewRelic %v.\nMessage: %s\n\n", r.Success, n.string)
	return nil
}

func main() {
	if err := Send(&Slack{"test Slack message"}); err != nil {
		// Handle errors better than me
		panic(err)
	}

	if err := Send(&NewRelic{"test NewRelic message"}); err != nil {
		// Handle errors better than me
		panic(err)
	}
}
