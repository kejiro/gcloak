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

type Keycloak interface {
	ServerInfo() (*representations.ServerInfo, error)
	Realm(name string) RealmResource
	CreateRealm(realm *representations.Realm) error
	Realms() ([]*representations.Realm, error)
}

func New(baseUrl string, httpClient *http.Client) Keycloak {
	return &keycloak{
		httpClient: httpClient,
		baseUrl:    baseUrl,
	}
}

type keycloak struct {
	baseUrl    string
	httpClient *http.Client
}

func (c *keycloak) composeUrl(paths ...string) string {
	return strings.Join(append([]string{c.baseUrl, "admin"}, paths...), "/")
}

func (c *keycloak) ServerInfo() (*representations.ServerInfo, error) {
	url := c.composeUrl("serverinfo")
	info := new(representations.ServerInfo)
	if err := c.doGet(url, info); err != nil {
		return nil, err
	}
	return info, nil
}

func (c *keycloak) Realm(name string) RealmResource {
	return &realmsResource{c, name}
}

func (c *keycloak) CreateRealm(realm *representations.Realm) error {
	url := c.composeUrl("realms")
	_, err := c.doCreate(url, realm)
	return err
}

func (c *keycloak) Realms() ([]*representations.Realm, error) {
	url := c.composeUrl("realms")
	realms := make([]*representations.Realm, 0)
	if err := c.doList(url, &realms); err != nil {
		return nil, err
	}
	return realms, nil
}

func (c *keycloak) doCreate(url string, obj interface{}) (string, error) {
	var reqBody *bytes.Reader
	if obj != nil {
		body, err := jsoniter.Marshal(obj)
		if err != nil {
			return "", err
		}
		reqBody = bytes.NewReader(body)
	}
	req, err := http.NewRequest(http.MethodPost, url, reqBody)
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
	if userUri == "" {
		return "", nil
	}
	paths := strings.Split(userUri, "/")
	if len(paths) == 0 {
		return "", ErrInvalidId
	}
	return paths[len(paths)-1], nil
}

func (c *keycloak) doGet(url string, obj interface{}) error {
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

func (c *keycloak) doUpdate(url string, obj interface{}) error {
	var reqBody *bytes.Reader
	if obj != nil {
		body, err := jsoniter.Marshal(obj)
		if err != nil {
			return err
		}
		reqBody = bytes.NewReader(body)
	}

	req, err := http.NewRequest(http.MethodPut, url, reqBody)
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

func (c *keycloak) doCount(url string) (int, error) {
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

func (c *keycloak) doList(url string, obj interface{}) error {
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

func (c *keycloak) doDelete(url string) error {
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

func (c *keycloak) do(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}
