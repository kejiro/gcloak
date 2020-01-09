package representations

type Role struct {
	Attributes  Map            `json:"attributes,omitempty"`
	ClientRole  *bool          `json:"clientRole,omitempty"`
	Composite   *bool          `json:"composite,omitempty"`
	Composites  RoleComposites `json:"composites,omitempty"`
	ContainerId string         `json:"containerId,omitempty"`
	Description string         `json:"description,omitempty"`
	Id          string         `json:"id,omitempty"`
	Name        string         `json:"name,omitempty"`
}

type RoleComposites struct {
	Client Map      `json:"client,omitempty"`
	Realm  []string `json:"realm,omitempty"`
}
