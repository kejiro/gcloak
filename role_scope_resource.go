package gcloak

import (
	"bytes"
	"errors"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/kejiro/gcloak/representations"
)

type RoleScopeResource interface {
	Add(roles []*representations.Role) error
	List() ([]*representations.Role, error)
	ListAvailable() ([]*representations.Role, error)
	ListEffective() ([]*representations.Role, error)
	Remove(roles []*representations.Role) error
}

type roleScopeResource struct {
	api
	roleType string
}

func (r *roleScopeResource) composeUrl(paths ...string) string {
	return r.api.composeUrl(append([]string{r.roleType}, paths...)...)
}

func (r *roleScopeResource) Add(roles []*representations.Role) error {
	url := r.composeUrl()
	_, err := r.doCreate(url, roles)
	return err
}

func (r *roleScopeResource) List() ([]*representations.Role, error) {
	url := r.composeUrl()
	roles := make([]*representations.Role, 0)
	if err := r.doGet(url, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleScopeResource) ListAvailable() ([]*representations.Role, error) {
	url := r.composeUrl("available")
	roles := make([]*representations.Role, 0)
	if err := r.doGet(url, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleScopeResource) ListEffective() ([]*representations.Role, error) {
	url := r.composeUrl("composite")
	roles := make([]*representations.Role, 0)
	if err := r.doGet(url, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleScopeResource) Remove(roles []*representations.Role) error {
	url := r.composeUrl()
	body, err := jsoniter.Marshal(roles)
	if err != nil {
		return err
	}
	reqBody := bytes.NewReader(body)
	req, err := http.NewRequest(http.MethodDelete, url, reqBody)
	if err != nil {
		return err
	}
	res, err := r.do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		return errors.New(res.Status)
	}
	return nil
}
