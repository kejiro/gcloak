package representations

type UserConsent struct {
	ClientId            string   `json:"clientId,omitempty"`
	CreatedDate         int64    `json:"createdDate,omitempty"`
	GrantedClientScopes []string `json:"grantedClientScopes,omitempty"`
	LastUpdatedDate     int64    `json:"lastUpdatedDate,omitempty"`
}
