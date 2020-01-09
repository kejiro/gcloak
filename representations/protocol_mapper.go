package representations

type ProtocolMapper struct {
	Config         Map    `json:"config,omitempty"`
	Id             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Protocol       string `json:"protocol,omitempty"`
	ProtocolMapper string `json:"protocolMapper,omitempty"`
}
