package issue

import (
	"fmt"
	"net/http"
)

type IssueModel struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	State string `json:"state"`
	Body  string `json:"body"`
}

func (c *Issue) ListIssue() ([]*IssueModel, error) {
	relativePath := "/issues"

	req, err := c.newRequest(http.MethodGet, relativePath, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	// send request
	var issueModel []*IssueModel
	statusCode, err := c.doRequest(req, &issueModel, nil)

	switch statusCode {
	case http.StatusOK:
		return issueModel, nil
	case http.StatusBadRequest:
		return nil, fmt.Errorf("StatusBadRequest")
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("StatusUnauthorized")
	case http.StatusForbidden:
		return nil, fmt.Errorf("StatusForbidden")
	case http.StatusNotFound:
		return nil, fmt.Errorf("StatusNotFound")
	default:
		fmt.Println(statusCode)
		return nil, err
	}
}
