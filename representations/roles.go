package representations

type Roles struct {
	Client Map    `json:"client,omitempty"`
	Realm  []Role `json:"realm,omitempty"`
}
