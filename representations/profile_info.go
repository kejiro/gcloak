package representations

type ProfileInfo struct {
	DisabledFeatures     []string `json:"disabledFeatures,omitempty"`
	ExperimentalFeatures []string `json:"experimentalFeatures,omitempty"`
	Name                 string   `json:"name,omitempty"`
	PreviewFeatures      []string `json:"previewFeatures,omitempty"`
}
