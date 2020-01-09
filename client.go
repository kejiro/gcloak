package gcloak

import (
	"errors"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/kejiro/gcloak/representations"
)

type Client interface {
	ServerInfo() (representations.ServerInfo, error)
}

func New(baseUrl string, httpClient *http.Client) Client {
	return &client{
		httpClient: httpClient,
		baseUrl:    baseUrl,
	}
}

type client struct {
	baseUrl    string
	httpClient *http.Client
}

func (c *client) ServerInfo() (representations.ServerInfo, error) {
	info := representations.ServerInfo{}
	res, err := c.httpClient.Get(c.baseUrl + "/admin/")
	if err != nil {
		return info, err
	}
	switch res.StatusCode {
	case http.StatusUnauthorized:
		return info, ErrAccessDenied
	case http.StatusNotFound:
		return info, ErrNotFound
	case http.StatusOK:
	// Do nothing
	default:
		return info, errors.New(res.Status)
	}

	dec := jsoniter.NewDecoder(res.Body)
	if err := dec.Decode(&info); err != nil {
		return info, err
	}
	return info, nil
}
