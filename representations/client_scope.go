package representations

type ClientScope struct {
	Attributes      Map              `json:"attributes,omitempty"`
	Description     string           `json:"description,omitempty"`
	Id              string           `json:"id,omitempty"`
	Name            string           `json:"name,omitempty"`
	Protocol        string           `json:"protocol,omitempty"`
	ProtocolMappers []ProtocolMapper `json:"protocolMappers,omitempty"`
}
