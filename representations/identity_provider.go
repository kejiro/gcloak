package representations

type IdentityProvider struct {
	AddReadTokenRoleOnCreate  *bool  `json:"addReadTokenRoleOnCreate,omitempty"`
	Alias                     string `json:"alias,omitempty"`
	Config                    Map    `json:"config,omitempty"`
	DisplayName               string `json:"displayName,omitempty"`
	Enabled                   *bool  `json:"enabled,omitempty"`
	FirstBrokerLoginFlowAlias string `json:"firstBrokerLoginFlowAlias,omitempty"`
	InternalId                string `json:"internalId,omitempty"`
	LinkOnly                  *bool  `json:"linkOnly,omitempty"`
	PostBrokerLoginFlowAlias  string `json:"postBrokerLoginFlowAlias,omitempty"`
	ProviderId                string `json:"providerId,omitempty"`
	StoreToken                string `json:"storeToken,omitempty"`
	TrustEmail                *bool  `json:"trustEmail,omitempty"`
}
