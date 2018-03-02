package hypertrack

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

const PROTOCOL = "https"
const HOST = "api.hypertrack.com"
const VERSION = "v1"
const TIMEOUT = 30

/*
Client to the HTTP API of Hypertrack.
*/
type Client struct {
	Key        string
	Secret     string
	HttpClient *http.Client
}

/*
There easiest way to configure the library is by creating a new `Hypertrack` instance:
	htClient := hypertrack.NewClient("pk_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "sk_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
*/
func NewClient(key, secret string) *Client {
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   TIMEOUT * time.Second,
			KeepAlive: TIMEOUT * time.Second,
			DualStack: true,
		}).DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       time.Second * 60,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * TIMEOUT,
	}

	return &Client{
		Key:        key,
		Secret:     secret,
		HttpClient: client,
	}
}

/*
Send request
*/
func (c *Client) request(method, url string, body []byte) ([]byte, error) {
	return request(c, method, url, body)
}

/*
Create User
*/
func (c *Client) CreateUser(name, phone, photo, lookupId, groupId string) (*User, error) {
	url := fmt.Sprintf("%s://%s/api/%s/users/", PROTOCOL, HOST, VERSION)

	request := make(map[string]interface{})
	if name != "" {
		request["name"] = name
	}
	if phone != "" {
		request["phone"] = phone
	}
	if photo != "" {
		request["photo"] = photo
	}
	if lookupId != "" {
		request["lookup_id"] = lookupId
	}
	if groupId != "" {
		request["group_id"] = groupId
	}

	requestPayload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	responsePayload, err := c.request("POST", url, requestPayload)
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(responsePayload, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

/*
Retrieve user
*/
func (c *Client) RetrieveUser(userId string) (*User, error) {
	url := fmt.Sprintf("%s://%s/api/%s/users/%s/", PROTOCOL, HOST, VERSION, userId)

	responsePayload, err := c.request("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(responsePayload, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

/*
Assign action to a user
*/
func (c *Client) AssignActionToUser(userId string, actionIds []string) (*User, error) {
	url := fmt.Sprintf("%s://%s/api/%s/users/%s/assign_actions/", PROTOCOL, HOST, VERSION, userId)

	request := make(map[string]interface{})
	request["action_ids"] = actionIds

	requestPayload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	responsePayload, err := c.request("POST", url, requestPayload)
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(responsePayload, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

/*
Create action
*/
func (c *Client) CreateActionUsingAddress(address, city, zipCode, country, lookupId string, actionType ActionType, scheduledAt *time.Time) (*Action, error) {
	url := fmt.Sprintf("%s://%s/api/%s/actions/", PROTOCOL, HOST, VERSION)

	expectedPlace := make(map[string]string)
	expectedPlace["address"] = address
	expectedPlace["city"] = city
	expectedPlace["zip_code"] = zipCode
	expectedPlace["country"] = country

	request := make(map[string]interface{})
	request["expected_place"] = expectedPlace
	if lookupId != "" {
		request["lookup_id"] = lookupId
	}
	if scheduledAt != nil {
		request["scheduled_at"] = scheduledAt
	}

	requestPayload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	responsePayload, err := c.request("POST", url, requestPayload)
	if err != nil {
		return nil, err
	}

	var action Action
	err = json.Unmarshal(responsePayload, &action)
	if err != nil {
		return nil, err
	}

	return &action, nil
}

/*
Complete action
*/
func (c *Client) CompleteAction(actionId string) (*Action, error) {
	url := fmt.Sprintf("%s://%s/api/%s/actions/%s/complete/", PROTOCOL, HOST, VERSION, actionId)

	request := make(map[string]interface{})
	// request["completion_time"] = time.Now()
	// request["completion_location"] =

	requestPayload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	responsePayload, err := c.request("POST", url, requestPayload)
	if err != nil {
		return nil, err
	}

	var action Action
	err = json.Unmarshal(responsePayload, &action)
	if err != nil {
		return nil, err
	}

	return &action, nil
}

/*
Cancel action
*/
func (c *Client) CancelAction(actionId string) (*Action, error) {
	url := fmt.Sprintf("%s://%s/api/%s/actions/%s/cancel/", PROTOCOL, HOST, VERSION, actionId)

	request := make(map[string]interface{})

	requestPayload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	responsePayload, err := c.request("POST", url, requestPayload)
	if err != nil {
		return nil, err
	}

	var action Action
	err = json.Unmarshal(responsePayload, &action)
	if err != nil {
		return nil, err
	}

	return &action, nil
}

/*
Retrieve action
*/
func (c *Client) RetrieveAction(actionId string) (*Action, error) {
	url := fmt.Sprintf("%s://%s/api/%s/actions/%s/", PROTOCOL, HOST, VERSION, actionId)

	responsePayload, err := c.request("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var action Action
	err = json.Unmarshal(responsePayload, &action)
	if err != nil {
		return nil, err
	}

	return &action, nil
}
