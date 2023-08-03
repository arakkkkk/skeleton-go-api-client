package issue

import (
  "go-api-client-sample/samplesdk/client"
	"io"
	"log"
	"net/http"
)

type Issue struct {
	*client.Client
}

func New(cfg *client.Config, logger *log.Logger) (*Issue, error) {
	c, err := client.New(cfg, logger)
	svc := &Issue{
		Client: c,
	}
	return svc, err
}

func (c *Issue) newRequest(method, relativePath string, queries, headers map[string]string, reqBody io.Reader) (*http.Request, error) {
	req, err := c.NewRequest(method, relativePath, queries, headers, reqBody)
	return req, err
}

func (c *Issue) doRequest(req *http.Request, respBody, respErr interface{}) (int, error) {
	statusCode, err := c.DoRequest(req, respBody, respErr)
	return statusCode, err
}
