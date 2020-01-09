package representations

type Group struct {
	Access      Map      `json:"access,omitempty"`
	Attributes  Map      `json:"attributes,omitempty"`
	ClientRoles Map      `json:"clientRoles,omitempty"`
	Id          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Path        string   `json:"path,omitempty"`
	RealmRoles  []string `json:"realmRoles,omitempty"`
	SubGroups   []Group  `json:"subGroups,omitempty"`
}
