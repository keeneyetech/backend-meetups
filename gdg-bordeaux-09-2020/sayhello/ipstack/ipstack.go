package ipstack

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IPStack struct{}

func New() *IPStack {
	return &IPStack{}
}

func (is IPStack) GetCountryCode(IP string) (string, error) {
	const ipstackURL = "http://api.ipstack.com/%s?access_key=40d1f026ac4207d3011dcffc31e4a4a6"
	resp, err := http.Get(fmt.Sprintf(ipstackURL, IP))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error contacting ipstack API: get code %d", resp.StatusCode)
	}
	// 2. Decode response to get location
	var location struct {
		CountryCode string `json:"country_code"`
	}
	err = json.NewDecoder(resp.Body).Decode(&location)
	return location.CountryCode, err
}
