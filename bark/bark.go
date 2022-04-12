package bark

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL     string
	Key         string
	PushOptions *PushOptions
}

type Options struct {
	Category          string `json:"category"`
	Level             string `json:"level,omitempty"`
	Badge             string `json:"badge,omitempty"`
	AutomaticallyCopy string `json:"automaticallyCopy,omitempty"`
	Copy              string `json:"code,omitempty"`
	Sound             string `json:"sound,omitempty"`
	Icon              string `json:"icon,omitempty"`
	Archive           string `json:"isArchive,omitempty"`
	Url               string `json:"url,omitempty"`
	Group             string `json:"group,omitempty"`
}

type PushOptions struct {
	Options
	Title     string `json:"title"`
	Body      string `json:"body"`
	DeviceKey string `json:"device_key"`
}

type PushResponse struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func NewClient(baseURL string, key string) *Client {
	c := &Client{BaseURL: baseURL, Key: key}
	return c
}

func (c *Client) Push(title string, body string, options *Options) (*PushResponse, error) {
	requestURL, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}
	pushOptions := &PushOptions{Title: title, Body: body, DeviceKey: c.Key, Options: *options}

	respBody, err := json.Marshal(pushOptions)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(requestURL.String(), "application/json", bytes.NewBuffer(respBody))
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	pushResp := &PushResponse{}
	err = json.Unmarshal(b, pushResp)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 || pushResp.Code != 200 {
		return pushResp, errors.New(pushResp.Message)
	}
	return pushResp, nil
}
