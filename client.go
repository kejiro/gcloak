package gcloak

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/kejiro/gcloak/representations"
)

type Client interface {
	ServerInfo() (representations.ServerInfo, error)
	Realm(name string) RealmResource
	CreateRealm(realm *representations.Realm) error
	ListRealms() ([]*representations.Realm, error)
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

func (c *client) composeUrl(paths ...string) string {
	return strings.Join(append([]string{c.baseUrl, "admin"}, paths...), "/")
}

func (c *client) ServerInfo() (representations.ServerInfo, error) {
	info := representations.ServerInfo{}
	req, err := http.NewRequest(http.MethodGet, c.baseUrl+"/admin/serverinfo", nil)
	if err != nil {
		return info, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := c.httpClient.Do(req)
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

func (c *client) Realm(name string) RealmResource {
	return &realmsResource{c, name}
}

func (c *client) CreateRealm(realm *representations.Realm) error {
	url := c.composeUrl("realms")
	_, err := c.doCreate(url, realm)
	return err
}

func (c *client) ListRealms() ([]*representations.Realm, error) {
	url := c.composeUrl("realms")
	realms := make([]*representations.Realm, 0)
	if err := c.doList(url, &realms); err != nil {
		return nil, err
	}
	return realms, nil
}

func (c *client) doCreate(url string, obj interface{}) (string, error) {
	reqBody, err := jsoniter.Marshal(obj)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusCreated {
		return "", errors.New(res.Status)
	}

	userUri := res.Header.Get("Location")
	paths := strings.Split(userUri, "/")
	if len(paths) == 0 {
		return "", ErrInvalidId
	}
	return paths[len(paths)-1], nil
}

func (c *client) doGet(url string, obj interface{}) error {
	req, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}
	if req.StatusCode != http.StatusOK {
		return errors.New(req.Status)
	}
	dec := jsoniter.NewDecoder(req.Body)
	return dec.Decode(obj)
}

func (c *client) doUpdate(url string, obj interface{}) error {
	reqBody, err := jsoniter.Marshal(obj)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		return errors.New(res.Status)
	}
	return nil
}

func (c *client) doCount(url string) (int, error) {
	res, err := c.httpClient.Get(url)
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, errors.New(res.Status)
	}

	body := make([]byte, res.ContentLength)

	if _, err := res.Body.Read(body); err != nil && err != io.EOF {
		return 0, err
	}
	cnt, err := strconv.ParseInt(string(body), 10, 0)
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}

func (c *client) doList(url string, obj interface{}) error {
	res, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	dec := jsoniter.NewDecoder(res.Body)
	return dec.Decode(obj)
}

func (c *client) doDelete(url string) error {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		return errors.New(res.Status)
	}
	return nil
}
