package representations

type Policy struct {
	Config           Map              `json:"config,omitempty"`
	DecisionStrategy DecisionStrategy `json:"decisionStrategy,omitempty"`
	Description      string           `json:"description,omitempty"`
	Id               string           `json:"id,omitempty"`
	Logic            Logic            `json:"logic,omitempty"`
	Name             string           `json:"name,omitempty"`
	Owner            string           `json:"owner,omitempty"`
	Policies         []string         `json:"policies,omitempty"`
	Resources        []string         `json:"resources,omitempty"`
	Scopes           []string         `json:"scopes,omitempty"`
	PolicyType       string           `json:"type,omitempty"`
}
