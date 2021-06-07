package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

type (
	// IP type is an example structure of IP information, along with more specifics that would be returned
	// from an IP authority.
	IP struct {
		// Inherit the net.IP data model
		IP net.IP

		// Additional attributes attached to the data that will allow us to
		// enrich the net.IP struct
		Info struct {
			AS          string  `json:"as"`
			City        string  `json:"city"`
			Country     string  `json:"country"`
			CountryCode string  `json:"countryCode"`
			ISP         string  `json:"isp"`
			Lat         float64 `json:"lat"`
			Lon         float64 `json:"lon"`
			Org         string  `json:"org"`
			Query       string  `json:"query"`
			Region      string  `json:"region"`
			RegionName  string  `json:"regionName"`
			Status      string  `json:"status"`
			Timezone    string  `json:"timezone"`
			Zip         string  `json:"zip"`
		}
	}
)

func main() {
	// Create a set of data using the API type.
	ipRecv := IP{
		IP: net.ParseIP("50.56.20.10"),
	}

	// Now that we have our data, we can access methods associated with that data using the
	// dcxReceiver as the receiver value.
	if err := ipRecv.Fetch(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(ipRecv.String())
}

// Fetch is a real-world example of a method that would require a certain set of data from its receiver
// to be able to execute the actions defined.
func (ip *IP) Fetch() error {
	req, err := http.NewRequest("GET", "http://ip-api.com/json/"+ip.IP.String(), nil)
	if err != nil {
		return err
	}

	doResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(doResp.Body).Decode(&ip.Info); err != nil {
		return fmt.Errorf("json.NewDecoder().Decode(%s)", err)
	}
	return nil
}

// String is another example of a method attached to the IP data type that is used as a simple helper to
// return some of the data points in a formatted string.
func (ip *IP) String() string {
	return fmt.Sprintf("[%s] as=%s city=%s region=%s", ip.IP.String(), ip.Info.AS, ip.Info.City, ip.Info.Region)
}
