package gcloak

import (
	"github.com/kejiro/gcloak/representations"
)

type ClientsResource interface {
	Create(role *representations.Client) (id string, err error)
	Get(id string) (*representations.Client, error)
	Update(id string, role *representations.Client) error
	Delete(id string) error
	Count() (int, error)
	List() ([]*representations.Client, error)
	Resource(id string) ClientResource
}

type clientsResource struct {
	*realmsResource
}

func (c *clientsResource) composeUrl(paths ...string) string {
	return c.realmsResource.composeUrl(append([]string{"clients"}, paths...)...)
}

func (c *clientsResource) Create(client *representations.Client) (id string, err error) {
	url := c.composeUrl()
	return c.doCreate(url, client)
}

func (c *clientsResource) Get(id string) (*representations.Client, error) {
	url := c.composeUrl(id)
	client := new(representations.Client)
	if err := c.doGet(url, client); err != nil {
		return nil, err
	}
	return client, nil
}

func (c *clientsResource) Update(id string, role *representations.Client) error {
	url := c.composeUrl(id)
	return c.doUpdate(url, role)
}

func (c *clientsResource) Count() (int, error) {
	url := c.composeUrl()
	return c.doCount(url)
}

func (c *clientsResource) List() ([]*representations.Client, error) {
	url := c.composeUrl()
	list := make([]*representations.Client, 0, 100)
	if err := c.doList(url, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *clientsResource) Resource(id string) ClientResource {
	return &clientResource{c, id}
}
