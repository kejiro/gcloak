package representations

type AuthenticationFlow struct {
	Alias                    string                          `json:"alias,omitempty"`
	AuthenticationExecutions []AuthenticationExecutionExport `json:"authenticationExecutions,omitempty"`
	BuiltIn                  *bool                           `json:"builtIn,omitempty"`
	Description              string                          `json:"description,omitempty"`
	Id                       string                          `json:"id,omitempty"`
	ProviderId               string                          `json:"providerId,omitempty"`
	TopLevel                 *bool                           `json:"topLevel,omitempty"`
}
