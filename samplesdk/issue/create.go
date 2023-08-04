package issue

import (
	"bytes"
	"encoding/json"
)

type CreateIssueInput struct {
	Title string `json:"title"`
	Assignees []string `json:"assignees"`
}

func (c *Issue) CreateIssue() (issueModel *IssueModel, err error) {
	relativePath := "/repos/arakkkkk/skeleton-go-api-client/issues"

	createIssueInput := &CreateIssueInput{
		Title: "sample",
	}
	createIssueInput.Assignees = append(createIssueInput.Assignees, "arakkkkk")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(createIssueInput)

	req, err := c.newRequest("POST", relativePath, nil, nil, b)
	if err != nil {
		return nil, err
	}

	// send request
	err = c.doRequest(req, &issueModel, nil)
	return
}
