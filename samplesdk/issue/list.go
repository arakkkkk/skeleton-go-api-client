package issue

import (
	"net/http"
)

type IssueModel struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	State string `json:"state"`
	Body  string `json:"body"`
}

type ListIssueQuery struct {
	PerPage  int `json:"per_page"`
}

func (c *Issue) ListIssue(listIssueQuery *ListIssueQuery) (issueModel []*IssueModel, err error) {
	relativePath := "/issues"

	queries := map[string]string {
		"per_page": "2",
	}

	req, err := c.newRequest(http.MethodGet, relativePath, queries, nil, nil)
	if err != nil {
		return nil, err
	}

	// send request
	err = c.doRequest(req, &issueModel, nil)
	return
}
