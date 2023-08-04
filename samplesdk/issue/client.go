package issue

import (
	"fmt"
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
	fmt.Println(req.URL)
	return req, err
}

func (c *Issue) doRequest(req *http.Request, respBody, respErr interface{}) (error) {
	statusCode, err := c.DoRequest(req, respBody, respErr)
	switch statusCode {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		return fmt.Errorf("StatusBadRequest")
	case http.StatusUnauthorized:
		return fmt.Errorf("StatusUnauthorized")
	case http.StatusForbidden:
		return fmt.Errorf("StatusForbidden")
	case http.StatusNotFound:
		return fmt.Errorf("StatusNotFound")
	default:
		fmt.Println(statusCode)
		return err
	}
}
