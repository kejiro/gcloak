package gcloak

import (
	"context"
	"os"

	"golang.org/x/oauth2/clientcredentials"
)

func setupClient() Keycloak {
	baseUrl := os.Getenv("BASE_URL")
	config := clientcredentials.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		TokenURL:     baseUrl + "/realms/master/protocol/openid-connect/token",
	}
	httpClient := config.Client(context.TODO())
	return New(baseUrl, httpClient)

}
