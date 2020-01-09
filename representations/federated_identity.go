package representations

type FederatedIdentity struct {
	IdentityProvider string `json:"identityProvider,omitempty"`
	UserId           string `json:"userId,omitempty"`
	UserName         string `json:"userName,omitempty"`
}
