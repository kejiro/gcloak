package gcloak

import (
	"github.com/kejiro/gcloak/representations"
)

type clientRoleScopeResource struct {
	api
	id string
}

func (u *clientRoleScopeResource) composeUrl(paths ...string) string {
	return u.api.composeUrl(append([]string{"clients", u.id}, paths...)...)
}

func (u *clientRoleScopeResource) Add(roles []*representations.Role) error {
	panic("implement me")
}

func (u *clientRoleScopeResource) List() ([]*representations.Role, error) {
	panic("implement me")
}

func (u *clientRoleScopeResource) ListAvailable() ([]*representations.Role, error) {
	panic("implement me")
}

func (u *clientRoleScopeResource) ListEffective() ([]*representations.Role, error) {
	panic("implement me")
}

func (u *clientRoleScopeResource) Remove(roles []*representations.Role) error {
	panic("implement me")
}
