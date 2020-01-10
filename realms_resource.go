package gcloak

// RealmsResource handles all the actions relating to one realm.
type RealmsResource interface {
	Users() UsersResource
	Clients() ClientsResource
	Roles() RolesResource
	Groups() GroupsResource
}

type realmsResource struct {
	*client
	realm string
}

func (r *realmsResource) composeUrl(paths ...string) string {
	return r.client.composeUrl(append([]string{"realms", r.realm}, paths...)...)
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
