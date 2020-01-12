package representations

type ClientMappings struct {
	Client   string `json:"client,omitempty"`
	Id       string `json:"id,omitempty"`
	Mappings []Role `json:"mappings,omitempty"`
}
