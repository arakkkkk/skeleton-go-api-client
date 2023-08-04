package issue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateIssueInput struct {
	Title string `json:"title"`
	Assignees string `json:"assignees"`
}

func (c *Issue) CreateIssue() ([]*IssueModel, error) {
	relativePath := "/repos/arakkkkk/go-api-client-sample/issues"

	createIssueInput := &CreateIssueInput{
		Title: "sample",
		Assignees: "[\"arakkkkk\"]",
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(createIssueInput)

	req, err := c.newRequest(http.MethodGet, relativePath, nil, nil, b)
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
