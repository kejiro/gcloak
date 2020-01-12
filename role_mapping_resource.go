package gcloak

import "github.com/kejiro/gcloak/representations"

type RoleMappingResource interface {
	Client(id string) RoleScopeResource
	All() *representations.Mappings
	Realm() RoleScopeResource
}

type roleMappingResource struct {
	api
}

func (u *roleMappingResource) composeUrl(paths ...string) string {
	return u.api.composeUrl(append([]string{"role-mappings"}, paths...)...)
}

func (u *roleMappingResource) Client(id string) RoleScopeResource {
	return &roleScopeResource{u, "clients/" + id}
}

func (u *roleMappingResource) All() *representations.Mappings {
	panic("implement me")
}

func (u *roleMappingResource) Realm() RoleScopeResource {
	return &roleScopeResource{u, "realm"}
}
