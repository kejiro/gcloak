package representations

type IdentityProviderMapper struct {
	Config                 Map    `json:"config,omitempty"`
	Id                     string `json:"id,omitempty"`
	IdentityProviderAlias  string `json:"identityProviderAlias,omitempty"`
	IdentityProviderMapper string `json:"identityProviderMapper,omitempty"`
	Name                   string `json:"name,omitempty"`
}
