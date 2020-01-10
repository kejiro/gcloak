package gcloak

import (
	"github.com/kejiro/gcloak/representations"
)

type UsersResource interface {
	Count() (int, error)
	Create(user *representations.User) (id string, err error)
	Delete(id string) error
	Get(id string) (*representations.User, error)
	List() ([]*representations.User, error)
	Search(username string) ([]*representations.User, error)
	Update(id string, user *representations.User) error
}

type usersResource struct {
	*realmsResource
}

func (u *usersResource) composeUrl(paths ...string) string {
	return u.realmsResource.composeUrl(append([]string{"users"}, paths...)...)
}

func (u *usersResource) Count() (int, error) {
	url := u.composeUrl("count")
	return u.doCount(url)
}

func (u *usersResource) Create(user *representations.User) (string, error) {
	url := u.composeUrl()
	return u.doCreate(url, user)
}

func (u *usersResource) Get(id string) (*representations.User, error) {
	url := u.composeUrl(id)
	user := new(representations.User)
	if err := u.doGet(url, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usersResource) List() ([]*representations.User, error) {
	url := u.composeUrl()
	users := make([]*representations.User, 0, 100)
	if err := u.doList(url, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *usersResource) Search(username string) ([]*representations.User, error) {
	panic("implement me")
}

func (u *usersResource) Update(id string, user *representations.User) error {
	url := u.composeUrl(id)
	return u.doUpdate(url, user)
}
