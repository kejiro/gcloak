package representations

type AuthenticatorConfig struct {
	HelpText   string           `json:"helpText,omitempty"`
	Name       string           `json:"name,omitempty"`
	Properties []ConfigProperty `json:"properties,omitempty"`
	ProviderId string           `json:"providerId,omitempty"`
}
