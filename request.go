package hypertrack

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/golang/glog"
)

// change timeout to time.Duration
func request(client *Client, method, url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+client.Secret)

	glog.Infoln("HYPERTRACK REQUEST", req)
	glog.Infoln("HYPERTRACK REQUEST METHOD", req.Method)

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return processResponse(resp)
}

func processResponse(response *http.Response) ([]byte, error) {
	responseBody, err := ioutil.ReadAll(response.Body)

	glog.Infoln("HYPERTRACK RESPONSE BODY: ", string(responseBody))

	if err != nil {
		return nil, err
	}

	if response.StatusCode == 200 {
		return responseBody, nil
	}
	message := fmt.Sprintf("Status Code: %s - %s", strconv.Itoa(response.StatusCode), string(responseBody))
	err = errors.New(message)
	return nil, err
}
