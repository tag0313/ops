package data

type Login struct {
	Uid      string
	Pubkaddr string
}

type LoginResult struct {
	Token string `json:"token"`
}
