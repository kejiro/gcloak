package representations

type MultiValue struct {
	Empty      bool    `json:"enabled,omitempty"`
	LoadFactor float64 `json:"loadFactor,omitempty"`
	Threshold  int32   `json:"threshold,omitempty"`
}

type MultiValuedHashMap map[string]MultiValue
