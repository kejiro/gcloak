package representations

type UserFederationMapper struct {
	Config                        Map    `json:"config,omitempty"`
	FederationMapperType          string `json:"federationMapperType,omitempty"`
	FederationProviderDisplayName string `json:"federationProviderDisplayName,omitempty"`
	Id                            string `json:"id,omitempty"`
	Name                          string `json:"name,omitempty"`
}
