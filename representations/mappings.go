package representations

type Mappings struct {
	Clients map[string]ClientMappings `json:"clientMappings,omitempty"`
	Realm   []Role                    `json:"realmMappings,omitempty"`
}
