package gcloak

import (
	"errors"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/kejiro/gcloak/representations"
)

type ClientResource interface {
	GenerateSecret() (*representations.Credential, error)
	ClientSecret() (*representations.Credential, error)
	ServiceAccountUser() (*representations.User, error)
	Roles() RolesResource
}

type clientResource struct {
	*clientsResource
	id string
}

func (c *clientResource) composeUrl(paths ...string) string {
	return c.clientsResource.composeUrl(append([]string{c.id}, paths...)...)
}

func (c *clientResource) ServiceAccountUser() (*representations.User, error) {
	url := c.composeUrl("service-account-user")
	user := new(representations.User)
	if err := c.doGet(url, &user); err != nil {
		return nil, err
	}
	return user, nil
}

func (c *clientResource) GenerateSecret() (*representations.Credential, error) {
	url := c.composeUrl("client-secret")

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		return nil, errors.New(res.Status)
	}

	creds := new(representations.Credential)
	dec := jsoniter.NewDecoder(res.Body)
	if err := dec.Decode(creds); err != nil {
		return nil, err
	}
	return creds, nil
}

func (c *clientResource) ClientSecret() (*representations.Credential, error) {
	url := c.composeUrl("client-secret")
	creds := new(representations.Credential)
	if err := c.doGet(url, creds); err != nil {
		return nil, err
	}
	return creds, nil
}

func (c *clientResource) Roles() RolesResource {
	return &clientRolesResource{c}
}
