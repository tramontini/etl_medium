package model

import "time"

type Ticker struct {
	ID             string           `json:"id"`
	Name           string           `json:"name"`
	Symbol         string           `json:"symbol"`
	Rank           uint8            `json:"rank"`
	FirtAppearedAt time.Time        `json:"first_data_at"`
	Quotes         map[string]Quote `json:"quotes"`
}
