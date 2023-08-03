package items

import (
	"go-api-client-sample/samplesdk/client"
	"io"
	"log"
	"net/http"
)

type Items struct {
	*client.Client
}

func New(rawBaseURL, token string, logger *log.Logger) (*Items, error) {
	c, err := client.New(rawBaseURL, token, logger)
	svc := &Items{
		Client: c,
	}
	return svc, err
}

func (c *Items) newRequest(method, relativePath string, queries, headers map[string]string, reqBody io.Reader) (*http.Request, error) {
	req, err := c.newRequest(method, relativePath, queries, headers, reqBody)
	return req, err
}

func (c *Items) doRequest(req *http.Request, respBody interface{}) (int, error) {
	statusCode, err := c.doRequest(req, respBody)
	return statusCode, err
}
