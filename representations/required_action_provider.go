package representations

type RequiredActionProvider struct {
	Alias         string `json:"alias,omitempty"`
	Config        Map    `json:"config,omitempty"`
	DefaultAction *bool  `json:"defaultAction,omitempty"`
	Enabled       *bool  `json:"enabled,omitempty"`
	Name          string `json:"name,omitempty"`
	Priority      int32  `json:"priority,omitempty"`
	ProviderId    string `json:"providerId,omitempty"`
}
