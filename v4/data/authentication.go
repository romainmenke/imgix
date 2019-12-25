package data

type Authentication struct {
	Authorized bool        `json:"authorized"`
	ClientID   interface{} `json:"clientId"`
	Mode       string      `json:"mode"`
	ModeTitle  string      `json:"modeTitle"`
	Tag        string      `json:"tag"`
	User       string      `json:"user"`
}
