package gcloak

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/kejiro/gcloak/representations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServerInfo(t *testing.T) {
	client := setupClient()

	info, err := client.ServerInfo()
	require.NoError(t, err)
	require.NotNil(t, info)
	require.Equal(t, "8.0.1", info.SystemInfo.Version)
}

func TestRealms(t *testing.T) {
	client := setupClient()

	realms, err := client.Realms()
	require.NoError(t, err)
	require.NotEmpty(t, realms)
}

func TestCreateRealm(t *testing.T) {
	baseUrl := os.Getenv("BASE_URL")
	client := setupClient().(*keycloak)
	err := client.doDelete(baseUrl + "/admin/realms/test-realm")
	assert.NoError(t, err)
	expected := &representations.Realm{Realm: "test-realm", Enabled: &representations.TRUE}
	require.NoError(t, client.CreateRealm(expected))
	<-time.After(50 * time.Millisecond) // Wait to allow Keycloak to do its thing.
	actual := new(representations.Realm)
	require.NoError(t, client.doGet(baseUrl+"/admin/realms/test-realm", actual))
	require.Equal(t, "test-realm", actual.Realm)
}

func TestHelpers(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		switch req.URL.Path {
		case "/get":
			res.WriteHeader(http.StatusOK)
			_, _ = res.Write([]byte(`{"yay": "boo"}`))
		}
	}))
	defer srv.Close()
	client := setupClient().(*keycloak)
	client.httpClient = srv.Client()

	actual := struct {
		Yay string `json:"yay,omitempty"`
	}{}
	err := client.doGet(srv.URL+"/get", &actual)
	require.NoError(t, err)
	require.Equal(t, "boo", actual.Yay)
}
