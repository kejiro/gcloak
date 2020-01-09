package representations

type MemoryInfo struct {
	Free           int64  `json:"free,omitempty"`
	FreeFormatted  string `json:"freeFormated,omitempty"`
	FreePercentage int64  `json:"freePercentage,omitempty"`
	Total          int64  `json:"total,omitempty"`
	TotalFormatted string `json:"totalFormated,omitempty"`
	Used           int64  `json:"used,omitempty"`
	UsedFormatted  string `json:"usedFormated,omitempty"`
}
