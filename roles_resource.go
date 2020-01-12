package gcloak

import "github.com/kejiro/gcloak/representations"

type RolesResource interface {
	Create(role *representations.Role) (id string, err error)
	Get(name string) (*representations.Role, error)
	Update(name string, role *representations.Role) error
	Delete(name string) error
	Count() (int, error)
	List() ([]*representations.Role, error)
}

type rolesResource struct {
	api
}

func (r *rolesResource) composeUrl(paths ...string) string {
	return r.api.composeUrl(append([]string{"roles"}, paths...)...)
}

func (r *rolesResource) Create(role *representations.Role) (id string, err error) {
	url := r.composeUrl()
	return r.doCreate(url, role)
}

func (r *rolesResource) Get(name string) (*representations.Role, error) {
	url := r.composeUrl(name)

	role := new(representations.Role)
	if err := r.doGet(url, role); err != nil {
		return nil, err
	}
	return role, nil
}

func (r *rolesResource) Update(name string, role *representations.Role) error {
	url := r.composeUrl(name)
	return r.doUpdate(url, role)
}

func (r *rolesResource) Count() (int, error) {
	url := r.composeUrl("count")
	return r.doCount(url)
}

func (r *rolesResource) List() ([]*representations.Role, error) {
	url := r.composeUrl()
	roles := make([]*representations.Role, 0, 100)
	if err := r.doList(url, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *rolesResource) Delete(name string) error {
	url := r.composeUrl(name)
	return r.doDelete(url)
}
