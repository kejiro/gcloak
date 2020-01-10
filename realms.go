package gcloak

import (
	"bytes"
	"errors"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/kejiro/gcloak/representations"
)

type Realms interface {
	List() ([]representations.Realm, error)
	Create(realm representations.Realm) error
	Get(name string) (representations.Realm, error)
	Update(name string, realm representations.Realm) error
	Delete(name string) error
}

func (c *client) List() ([]representations.Realm, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseUrl+"/admin/realms", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	switch res.StatusCode {
	case http.StatusUnauthorized:
		return nil, ErrAccessDenied
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusOK:
	// Do nothing
	default:
		return nil, errors.New(res.Status)
	}

	dec := jsoniter.NewDecoder(res.Body)

	list := make([]representations.Realm, 0)
	if err := dec.Decode(&list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *client) Create(realm representations.Realm) error {
	body, err := jsoniter.Marshal(realm)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, c.baseUrl+"/admin/realms", bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusCreated {
		return errors.New(res.Status)
	}
	return nil
}

func (c *client) Get(name string) (representations.Realm, error) {
	realm := representations.Realm{}
	res, err := c.httpClient.Get(c.baseUrl + "/admin/realms/" + name)
	if err != nil {
		return realm, err
	}
	switch res.StatusCode {
	case http.StatusNotFound:
		return realm, ErrNotFound
	case http.StatusUnauthorized:
		return realm, ErrAccessDenied
	case http.StatusForbidden:
		return realm, ErrAccessDenied
	case http.StatusOK:
	// Do nothing
	default:
		return realm, errors.New(res.Status)
	}

	dec := jsoniter.NewDecoder(res.Body)
	if err := dec.Decode(&realm); err != nil {
		return realm, err
	}
	return realm, nil
}

func (c *client) Update(name string, realm representations.Realm) error {
	body, err := jsoniter.Marshal(realm)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, c.baseUrl+"/admin/realms/"+name, bytes.NewReader(body))
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

func (c *client) Delete(name string) error {
	req, err := http.NewRequest(http.MethodDelete, c.baseUrl+"/admin/realms/"+name, nil)
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
