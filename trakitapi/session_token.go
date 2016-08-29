package trakitapi

type SessionToken struct {
	Token      string `json:"token"`
	Tenant     string `json:"tenant"`
	Expiration string `json:"expiration"`
	Level      int    `json:"level"`
}

func (st *SessionToken) IsExpired() bool {
	return false
}

func (st *SessionToken) GetUser() string {
	return "user"
}
