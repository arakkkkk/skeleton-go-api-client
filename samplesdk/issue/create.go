package issue

// curl -L \
//   -X POST \
//   -H "Accept: application/vnd.github+json" \
//   -H "Authorization: Bearer <YOUR-TOKEN>" \
//   -H "X-GitHub-Api-Version: 2022-11-28" \
//   https://api.github.com/repos/OWNER/REPO/issues \
//   -d '{"title":"Found a bug","body":"I'\''m having a problem with this.","assignees":["octocat"],"milestone":1,"labels":["bug"]}'

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateIssueInput struct {
	Title string `json:"title"`
}

func (c *Issue) CreateIssue() ([]*IssueModel, error) {
	relativePath := "/issues"

	createIssueInput := &CreateIssueInput{
		Title: "sample",
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
