package representations

type Resource struct {
	Id                 string   `json:"id,omitempty"`
	Attributes         Map      `json:"attributes,omitempty"`
	DisplayName        string   `json:"displayName,omitempty"`
	IconUri            string   `json:"icon_uri,omitempty"`
	Name               string   `json:"name,omitempty"`
	OwnerManagedAccess *bool    `json:"ownerManagedAccess,omitempty"`
	Scopes             []Scope  `json:"scopes,omitempty"`
	ResourceType       string   `json:"resourceType,omitempty"`
	Uris               []string `json:"uris,omitempty"`
}
