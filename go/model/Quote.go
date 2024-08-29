package model

import "time"

type Quote struct {
	Price         float32   `json:"price"`
	LastUpdatedAt time.Time `json:"ath_date"`
}
