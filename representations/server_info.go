package representations

type ServerInfo struct {
	BuiltinProtocolMappers Map                  `json:"builtinProtocolMappers,omitempty"`
	ClientImporters        []Map                `json:"clientImporters,omitempty"`
	ClientInstallations    Map                  `json:"clientInstallations,omitempty"`
	ComponentTypes         Map                  `json:"componentTypes,omitempty"`
	Enums                  Map                  `json:"enums,omitempty"`
	IdentityProviders      []Map                `json:"identityProviders,omitempty"`
	MemoryInfo             MemoryInfo           `json:"memoryInfo,omitempty"`
	PasswordPolicies       []PasswordPolicyType `json:"passwordPolicies,omitempty"`
	ProfileInfo            ProfileInfo          `json:"profileInfo,omitempty"`
	ProtocolMapperTypes    Map                  `json:"protocolMapperTypes,omitempty"`
	Providers              Map                  `json:"providers,omitempty"`
	SocialProviders        []Map                `json:"socialProviders,omitempty"`
	SystemInfo             SystemInfo           `json:"systemInfo,omitempty"`
	Themes                 Map                  `json:"themes,omitempty"`
}
