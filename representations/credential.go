package representations

type Credential struct {
	CreatedDate    int64  `json:"createdDate,omitempty"`
	CredentialData string `json:"credentialData,omitempty"`
	Id             string `json:"id,omitempty"`
	Priority       int32  `json:"priority,omitempty"`
	SecretData     string `json:"secretData,omitempty"`
	Temporary      *bool  `json:"temporary,omitempty"`
	CredentialType string `json:"type,omitempty"`
	UserLabel      string `json:"userLabel,omitempty"`
	Value          string `json:"value,omitempty"`
}
