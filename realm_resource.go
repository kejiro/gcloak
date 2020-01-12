package gcloak

// RealmResource handles all the actions relating to one realm.
type RealmResource interface {
	Users() UsersResource
	Clients() ClientsResource
	Roles() RolesResource
	Groups() GroupsResource
}

type realmsResource struct {
	api
	realm string
}

func (r *realmsResource) composeUrl(paths ...string) string {
	return r.api.composeUrl(append([]string{"realms", r.realm}, paths...)...)
}

func (r *realmsResource) Users() UsersResource {
	return &usersResource{r}
}

func (r *realmsResource) Clients() ClientsResource {
	return &clientsResource{r}
}

func (r *realmsResource) Roles() RolesResource {
	return &rolesResource{r}
}

func (r *realmsResource) Groups() GroupsResource {
	return &groupsResource{r}
}
