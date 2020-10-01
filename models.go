package main

import "gorm.io/gorm"

type PaymentDetail struct {
	gorm.Model
	UrlPath string `json:"urlPath"`
	PaymentPointer string `json:"paymentPointer"`
	RequestID string `json:"requestId"`
	Amount string `json:"amount"`
	AssetCode string `json:"assetCode"`
	AssetScale int `json:"assetScale"`
	Receipt string `json:"receipt"`
}
