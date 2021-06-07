package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Github type is an example of a specific type of data (only instance defined in this example), that
// is tied to a certain set of actions.
//
// In this scenario, a data-set of the Github type has a method of Contributors exposed through the
// receiver.
type Github struct {
	Instance string
	Owner    string
}

// Contributor type will define the data returned from the Github API
// and will need to be used to decode the response value.
//
// HINT: Use the data model shown by the API docs to build your response type
//  - https://docs.github.com/en/rest/reference/repos#list-repository-contributors
//
// Don't forget to use struct tags
//  - https://godoc.org/encoding/json#Marshal
type Contributor struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	Contributions     int    `json:"contributions"`
}

func main() {
	// Create a set of data using the Github type.
	gitReceiver := Github{
		Instance: "https://api.github.com",
		Owner:    "josh5276",
	}

	repoName := "inventory-check"

	// Now that we have our data, we can access methods associated with that data using the
	// gitReceiver as the receiver value.
	r, err := gitReceiver.Contributors(repoName)
	if err != nil {
		log.Fatal(err)
	}

	// Use fmt.Println or another print method to output each contributor to the project,
	// along with how many contributions they have made.
	//
	// SAMPLE OUTPUT: "user josh5276 has made 14 contributions to the go-course repo"
	for _, d := range r {
		fmt.Printf("user %s has made %d contributions to the %s repo\n", d.Login, d.Contributions, repoName)
	}
}

// Contributors method returns a list of contributors to the specified repository and
// sorts them by the number of commits per contributor in descending order.
//
// See: https://docs.github.com/en/rest/reference/repos#list-repository-contributors
func (c Github) Contributors(repo string) ([]Contributor, error) {
	// Define our response data here. This will create a new instance of the
	// Contributor type with it's empty value.
	var data []Contributor

	// Build the URL that we will need to call to Github to get the information on repo contributors
	//
	// HINT: Full URL example https://api.github.com/josh5276/go-course/contributors
	reqURL := fmt.Sprintf("%s/repos/%s/%s/contributors", c.Instance, c.Owner, repo)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return data, err
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")
	doResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return data, err
	}

	if err := json.NewDecoder(doResp.Body).Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}
