package items

import (
	"errors"
	"fmt"
	"net/http"
)

type Item struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	LikesCount int    `json:"likes_count"`
}

func (c *Items) GetUserItems() ([]*Item, error) {
	relativePath := "/api/all"
	req, err := c.newRequest(http.MethodGet, relativePath, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	// send request
	var items []*Item
	code, err := c.doRequest(req, &items)

	switch code {
	case http.StatusOK:
		return items, nil
	case http.StatusBadRequest:
		return nil, errors.New("bad request. some parameters may be invalid")
	case http.StatusNotFound:
		return nil, fmt.Errorf("not found. user with id may not exist")
	default:
		return nil, errors.New("unexpected error")
	}
}


