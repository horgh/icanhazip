// Package icanhazip looks up your IP by asking icanhazip.com.
package icanhazip

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

// Lookup makes a request to icanhazip.com and retrieves our current IP.
func Lookup() (net.IP, error) {
	req, err := http.NewRequest("GET", "https://icanhazip.com", nil)
	if err != nil {
		return nil, fmt.Errorf("Unable to create request: %s", err)
	}

	client := &http.Client{}
	client.Timeout = time.Duration(60 * time.Second)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Request problem: %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	err2 := resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Unable to read body: %s", err)
	}
	if err2 != nil {
		return nil, fmt.Errorf("Problem closing body: %s", err2)
	}

	rawIP := strings.TrimSpace(string(body))

	ip := net.ParseIP(rawIP)
	if ip == nil {
		return nil, fmt.Errorf("Invalid IP received: %s", body)
	}

	return ip, nil
}
