package representations

type Session struct {
	Clients    Map    `json:"clients,omitempty"`
	Id         string `json:"id,omitempty"`
	IpAddress  string `json:"ipAddress,omitempty"`
	LastAccess int64  `json:"lastAccess,omitempty"`
	Start      int64  `json:"start,omitempty"`
	UserId     string `json:"userId,omitempty"`
	Username   string `json:"username,omitempty"`
}
