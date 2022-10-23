package models

type ListChanRequest struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}

type ListSessionRequest struct {
	Running bool `json:"running"`
	Limit   int  `json:"limit"`
	Skip    int  `json:"skip"`
}
