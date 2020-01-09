package representations

type PasswordPolicyType struct {
	ConfigType        string `json:"configType,omitempty"`
	DefaultValue      string `json:"defaultValue,omitempty"`
	DisplayName       string `json:"displayName,omitempty"`
	Id                string `json:"id,omitempty"`
	MultipleSupported *bool  `json:"multipleSupported,omitempty"`
}
