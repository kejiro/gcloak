package representations

type Scope struct {
	DisplayName string     `json:"displayName,omitempty"`
	IconUri     string     `json:"iconUri,omitempty"`
	Id          string     `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Policies    []Policy   `json:"policies,omitempty"`
	Resources   []Resource `json:"resources,omitempty"`
}
