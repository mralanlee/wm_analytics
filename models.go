package main

import (
	"time"
)

type PaymentDetail struct {
	ID uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UrlPath string `json:"urlPath" gorm:"index"`
	PaymentPointer string `json:"paymentPointer"`
	RequestID string `json:"requestId"`
	Amount string `json:"amount"`
	AssetCode string `json:"assetCode"`
	AssetScale int `json:"assetScale"`
	Receipt string `json:"receipt"`
}
