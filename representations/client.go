package representations

type Client struct {
	Access                             Map              `json:"access,omitempty"`
	AdminUrl                           string           `json:"adminUrl,omitempty"`
	Attributes                         Map              `json:"attributes,omitempty"`
	AuthenticationFlowBindingOverrides Map              `json:"authenticationFlowBindingOverrides,omitempty"`
	AuthorizationServicesEnabled       *bool            `json:"authorizationServicesEnabled,omitempty"`
	AuthorizationSettings              *ResourceServer  `json:"authorization_settings,omitempty"`
	BaseUrl                            string           `json:"baseUrl,omitempty"`
	BearerOnly                         *bool            `json:"bearerOnly,omitempty"`
	ClientAuthenticatorType            string           `json:"clientAuthenticatorType,omitempty"`
	ClientId                           string           `json:"clientId,omitempty"`
	ConsentRequired                    *bool            `json:"consentRequired,omitempty"`
	DefaultClientScopes                []string         `json:"defaultClientScopes,omitempty"`
	DefaultRoles                       []string         `json:"defaultRoles,omitempty"`
	Description                        string           `json:"description,omitempty"`
	DirectAccessGrantsEnabled          *bool            `json:"directAccessGrantsEnabled,omitempty"`
	Enabled                            *bool            `json:"enabled,omitempty"`
	FrontchannelLogout                 *bool            `json:"frontchannelLogout,omitempty"`
	FullScopeAllowed                   *bool            `json:"fullScopeAllowed,omitempty"`
	Id                                 string           `json:"id,omitempty"`
	ImplicitFlowEnabled                *bool            `json:"implicitFlowEnabled,omitempty"`
	Name                               string           `json:"name,omitempty"`
	NodeReRegistrationTimeout          int32            `json:"nodeReRegistrationTimeout,omitempty"`
	NotBefore                          int32            `json:"notBefore,omitempty"`
	OptionalClientScopes               []string         `json:"optionalClientScopes,omitempty"`
	Origin                             string           `json:"origin,omitempty"`
	Protocol                           string           `json:"protocol,omitempty"`
	ProtocolMappers                    []ProtocolMapper `json:"protocolMappers,omitempty"`
	PublicClient                       *bool            `json:"publicClient,omitempty"`
	RedirectUris                       []string         `json:"redirectUris,omitempty"`
	RegisteredNodes                    Map              `json:"registeredNodes,omitempty"`
	RegistrationAccessToken            string           `json:"registrationAccessToken,omitempty"`
	RootUrl                            string           `json:"rootUrl,omitempty"`
	Secret                             string           `json:"secret,omitempty"`
	ServiceAccountsEnabled             *bool            `json:"serviceAccountsEnabled,omitempty"`
	StandardFlowEnabled                *bool            `json:"standardFlowEnabled,omitempty"`
	SurrogateAuthRequired              *bool            `json:"surrogateAuthRequired,omitempty"`
	WebOrigins                         []string         `json:"webOrigins,omitempty"`
}
