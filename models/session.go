package models

type CreateSession struct {
	ChanID uint64                 `json:"chan_id"`
	IsTest bool                   `json:"is_test"`
	Data   map[string]interface{} `json:"data"`
}
