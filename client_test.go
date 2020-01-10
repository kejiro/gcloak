package gcloak

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/oauth2/clientcredentials"
)

type ClientSuite struct {
	suite.Suite

	baseUrl      string
	clientId     string
	clientSecret string
	client       Client
}

func TestClient(t *testing.T) {
	suite.Run(t, &ClientSuite{})
}

func (s *ClientSuite) SetupSuite() {
	s.baseUrl = os.Getenv("BASE_URL")
	s.clientId = os.Getenv("CLIENT_ID")
	s.clientSecret = os.Getenv("CLIENT_SECRET")
}

func (s *ClientSuite) SetupTest() {
	config := clientcredentials.Config{
		ClientID:     s.clientId,
		ClientSecret: s.clientSecret,
		TokenURL:     s.baseUrl + "/realms/master/protocol/openid-connect/token",
	}
	httpClient := config.Client(context.TODO())

	s.client = New(s.baseUrl, httpClient)
}

func (s *ClientSuite) TestClientList() {
	list, err := s.client.Realms("master").Clients().List()
	s.Require().NoError(err)
	s.Require().NotEmpty(list)
}
