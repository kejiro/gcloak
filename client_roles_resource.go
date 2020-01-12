package gcloak

import (
	"github.com/kejiro/gcloak/representations"
)

type clientRolesResource struct {
	api
}

func (c *clientRolesResource) composeUrl(paths ...string) string {
	return c.api.composeUrl(append([]string{"roles"}, paths...)...)
}

func (c *clientRolesResource) Create(role *representations.Role) (id string, err error) {
	url := c.composeUrl()
	return c.doCreate(url, role)
}

func (c *clientRolesResource) Get(name string) (*representations.Role, error) {
	url := c.composeUrl(name)
	role := new(representations.Role)
	if err := c.doGet(url, role); err != nil {
		return nil, err
	}
	return role, nil
}

func (c *clientRolesResource) Update(name string, role *representations.Role) error {
	url := c.composeUrl(name)
	return c.doUpdate(url, role)
}

func (c *clientRolesResource) Delete(name string) error {
	url := c.composeUrl(name)
	return c.doDelete(url)
}

func (c *clientRolesResource) Count() (int, error) {
	url := c.composeUrl()
	return c.doCount(url)
}

func (c *clientRolesResource) List() ([]*representations.Role, error) {
	url := c.composeUrl()
	roles := make([]*representations.Role, 0)
	if err := c.doList(url, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}
