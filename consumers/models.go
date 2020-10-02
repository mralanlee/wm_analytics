package consumers 

import (
	"log"
	"math"
	"strconv"
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
	AmountValue float64
}

func (p *PaymentDetail) transform() {
	amount, err := strconv.ParseFloat(p.Amount, 32)
	scale := math.Pow10(p.AssetScale)

	if err != nil {
		log.Fatal(err)
		return
	}

	p.AmountValue	= amount / scale
}
