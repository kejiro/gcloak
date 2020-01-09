package representations

type User struct {
	Access                     Map                 `json:"access,omitempty"`
	Attributes                 Map                 `json:"attributes,omitempty"`
	ClientConsents             []UserConsent       `json:"clientConsents,omitempty"`
	ClientRoles                Map                 `json:"clientRoles,omitempty"`
	CreatedTimestamp           int64               `json:"createdTimestamp,omitempty"`
	Credentials                []Credential        `json:"credentials,omitempty"`
	DisableableCredentialTypes []string            `json:"disableableCredentialTypes,omitempty"`
	Email                      string              `json:"email,omitempty"`
	EmailVerified              *bool               `json:"emailVerified,omitempty"`
	Enabled                    *bool               `json:"enabled,omitempty"`
	FederatedIdentities        []FederatedIdentity `json:"federatedIdentities,omitempty"`
	FederationLink             string              `json:"federationLink,omitempty"`
	FirstName                  string              `json:"firstName,omitempty"`
	Groups                     []string            `json:"groups,omitempty"`
	Id                         string              `json:"id,omitempty"`
	LastName                   string              `json:"lastName,omitempty"`
	NotBefore                  int32               `json:"notBefore,omitempty"`
	Origin                     string              `json:"origin,omitempty"`
	RealmRoles                 []string            `json:"realmRoles,omitempty"`
	RequiredActions            []string            `json:"requiredActions,omitempty"`
	Self                       string              `json:"self,omitempty"`
	ServiceAccountClientId     string              `json:"serviceAccountClientId,omitempty"`
	Username                   string              `json:"username,omitempty"`
}
