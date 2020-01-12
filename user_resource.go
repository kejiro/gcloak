package gcloak

import (
	"github.com/kejiro/gcloak/representations"
)

type UserResource interface {
	Roles() RoleMappingResource
	Groups() ([]*representations.Group, error)
	ResetPassword(credential *representations.Credential) error
	// Send an email with a reset password link to the user
	SendVerifyEmail() error
	Update(user *representations.User) error
	Logout() error
	Impersonate()
	JoinGroup(groupId string) error
	LeaveGroup(groupId string) error
	Sessions() ([]*representations.Session, error)
}

type userResource struct {
	api
	id string
}

func (u *userResource) composeUrl(paths ...string) string {
	return u.api.composeUrl(append([]string{u.id}, paths...)...)
}

func (u *userResource) Roles() RoleMappingResource {
	return &roleMappingResource{u}
}

func (u *userResource) Update(user *representations.User) error {
	url := u.composeUrl()
	return u.doUpdate(url, user)
}

func (u *userResource) Groups() ([]*representations.Group, error) {
	url := u.composeUrl("groups")
	groups := make([]*representations.Group, 0)
	if err := u.doGet(url, &groups); err != nil {
		return nil, err
	}
	return groups, nil
}

func (u *userResource) ResetPassword(credential *representations.Credential) error {
	url := u.composeUrl("reset-password")
	return u.doUpdate(url, credential)
}

func (u *userResource) SendVerifyEmail() error {
	url := u.composeUrl("send-verify-email")
	return u.doUpdate(url, nil)
}

func (u *userResource) Logout() error {
	url := u.composeUrl("logout")
	_, err := u.doCreate(url, nil)
	return err
}

func (u *userResource) Impersonate() {
	panic("implement me")
}

func (u *userResource) JoinGroup(groupId string) error {
	url := u.composeUrl("groups", groupId)
	return u.doUpdate(url, nil)
}

func (u *userResource) LeaveGroup(groupId string) error {
	url := u.composeUrl("groups", groupId)
	return u.doDelete(url)
}

func (u *userResource) Sessions() ([]*representations.Session, error) {
	url := u.composeUrl("sessions")
	sessions := make([]*representations.Session, 0)
	if err := u.doGet(url, &sessions); err != nil {
		return nil, err
	}
	return sessions, nil
}
