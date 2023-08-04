package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	Token      string
	Logger     *log.Logger
}

func New(cfg *Config, logger *log.Logger) (*Client, error) {
	baseURL, err := url.Parse(cfg.EndPoint)
	if err != nil {
		return nil, err
	}

	if logger == nil {
		logger = log.New(os.Stderr, "[LOG]", log.LstdFlags)
	}

	return &Client{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
		Token:      cfg.ApiToken,
		Logger:     logger,
	}, nil
}

func (c *Client) NewRequest(method, relativePath string, queries, headers map[string]string, reqBody io.Reader) (*http.Request, error) {
	// set path
	reqURL, err := c.BaseURL.Parse(relativePath)

	// set query
	if queries != nil {
		queries["key"] = c.Token
		q := reqURL.Query()
		for k, v := range queries {
			q.Add(k, v)
		}
		reqURL.RawQuery = q.Encode()
	}

	// instantiate request
	req, err := http.NewRequest(method, reqURL.String(), reqBody)
	if err != nil {
		return nil, err
	}

	// set header
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	return req, nil
}

func (c *Client) DoRequest(req *http.Request, respBody, respErr interface{}) (int, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(body))
	fmt.Print()

	if resp.StatusCode < 200 || 300 <= resp.StatusCode {
		json.NewDecoder(resp.Body).Decode(&respErr)
		return resp.StatusCode, nil
	}

	json.Unmarshal(body, &respBody)

	return resp.StatusCode, nil
}
