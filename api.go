package gcloak

import "net/http"

type api interface {
	composeUrl(paths ...string) string
	doCreate(url string, obj interface{}) (string, error)
	doGet(url string, obj interface{}) error
	doUpdate(url string, obj interface{}) error
	doCount(url string) (int, error)
	doList(url string, obj interface{}) error
	doDelete(url string) error
	do(req *http.Request) (*http.Response, error)
}
