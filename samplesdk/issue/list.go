package issue

import (
	"net/http"
	"strconv"
)

type IssueModel struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	State string `json:"state"`
	Body  string `json:"body"`
}

type ListIssueQuery struct {
	PerPage int
}

func (q *ListIssueQuery) toMap() map[string]string {
	return map[string]string{
		"per_page": strconv.Itoa(q.PerPage),
	}
}

func (c *Issue) ListIssue(listIssueQuery *ListIssueQuery) (issueModel []*IssueModel, err error) {
	relativePath := "/issues"

	queries := listIssueQuery.toMap()

	req, err := c.newRequest(http.MethodGet, relativePath, queries, nil, nil)
	if err != nil {
		return nil, err
	}

	// send request
	err = c.doRequest(req, &issueModel, nil)
	return
}
