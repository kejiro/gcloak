package representations

type UserFederationProvider struct {
	ChangedSyncPeriod int32  `json:"changedSyncPeriod,omitempty"`
	Config            Map    `json:"config,omitempty"`
	DisplayName       string `json:"displayName,omitempty"`
	FullSyncPeriod    int32  `json:"fullSyncPeriod,omitempty"`
	Id                string `json:"id,omitempty"`
	LastSync          int32  `json:"lastSync,omitempty"`
	Priority          int32  `json:"priority,omitempty"`
	ProviderName      string `json:"providerName,omitempty"`
}
