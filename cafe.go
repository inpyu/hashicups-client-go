package hashicups

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetCafes() ([]Cafe, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/cafes", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	cafes := []Cafe{}
	err = json.Unmarshal(body, &cafes)
	if err != nil {
		return nil, err
	}

	return cafes, nil
}

func (c *Client) GetCafe(cafeID string) (*Cafe, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/cafes/%s", c.HostURL, cafeID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	cafe := Cafe{}
	err = json.Unmarshal(body, &cafe)
	if err != nil {
		return nil, err
	}

	return &cafe, nil
}

func (c *Client) CreateCafe(cafes []Cafe) (*Cafe, error) {
	rb, err := json.Marshal(cafes)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/cafes", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	cafe := Cafe{}
	err = json.Unmarshal(body, &cafe)
	if err != nil {
		return nil, err
	}

	return &cafe, nil
}

func (c *Client) UpdateCafe(cafeID string, cafes []Cafe) (*Cafe, error) {
	rb, err := json.Marshal(cafes)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/cafes/%s", c.HostURL, cafeID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	cafe := Cafe{}
	err = json.Unmarshal(body, &cafe)
	if err != nil {
		return nil, err
	}

	return &cafe, nil
}

func (c *Client) DeleteCafe(cafeID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/cafes/%s", c.HostURL, cafeID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return err
	}

	if string(body) != "Deleted cafe" {
		return errors.New(string(body))
	}

	return nil
}
