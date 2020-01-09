package representations

type ConfigProperty struct {
	DefaultValue       Object   `json:"defaultValue,omitempty"`
	HelpText           string   `json:"helpText,omitempty"`
	Label              string   `json:"label,omitempty"`
	Name               string   `json:"name,omitempty"`
	Options            []string `json:"options,omitempty"`
	Secret             *bool    `json:"secret,omitempty"`
	ConfigPropertyType string   `json:"type,omitempty"`
}
