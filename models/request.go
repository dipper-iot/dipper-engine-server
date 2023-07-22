package models

type ListChanRequest struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}

type ListSessionRequest struct {
	Infinity *bool `json:"infinity"`
	Limit    int   `json:"limit"`
	Skip     int   `json:"skip"`
}
