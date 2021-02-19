package bark

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	BaseURL     string
	Key         string
	PushOptions *PushOptions
}

type PushOptions struct {
	Category    string
	Copy        string
	DirectedURL string
	AutoCopy    bool
	Archive     bool
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

func (c *Client) Push(title string, body string, options *PushOptions) (*PushResponse, error) {
	requestURL, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}
	requestURL.Path = path.Join(requestURL.Path, c.Key, title, body)

	if options != nil {
		q := &url.Values{}
		if options.Copy != "" {
			q.Add("copy", options.Copy)
		}
		if options.DirectedURL != "" {
			q.Add("url", options.DirectedURL)
		}
		if options.AutoCopy {
			q.Add("automaticallyCopy", "1")
		}
		if options.Archive {
			q.Add("isArchive", "1")
		}
		requestURL.RawQuery = q.Encode()
	}

	resp, err := http.Get(requestURL.String())
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
