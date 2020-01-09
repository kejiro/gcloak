package representations

type ResourceServer struct {
	AllowRemoteResourceManagement *bool                 `json:"allowRemoteResourceManagement,omitempty"`
	ClientId                      string                `json:"clientId,omitempty"`
	DecisionStrategy              DecisionStrategy      `json:"decisionStrategy,omitempty"`
	Id                            string                `json:"id,omitempty"`
	Name                          string                `json:"name,omitempty"`
	Policies                      []Policy              `json:"policies,omitempty"`
	PolicyEnforcementMode         PolicyEnforcementMode `json:"policyEnforcementMode,omitempty"`
	Resources                     []Resource            `json:"resources,omitempty"`
	Scopes                        []Scope               `json:"scopes,omitempty"`
}
