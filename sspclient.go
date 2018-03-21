package sspclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var baseURL = "http://ssp.news-host.pw"

// OverrideBaseURL ...
func OverrideBaseURL(url string) {
	baseURL = url
}

// Sell ...
func Sell(request *Request) error {
	url := fmt.Sprintf(baseURL + "/imp/")

	j, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if 202 != resp.StatusCode {
		return fmt.Errorf("%s", body)
	}

	return err
}
