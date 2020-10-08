package consumers

import (
	"log"
	"math"
	"strconv"
	"time"
)

type PaymentDetail struct {
	ID             uint      `gorm:"primaryKey"`
	ReceivedAt 		 time.Time `json:"receivedAt"`
	UrlPath        string    `json:"urlPath" gorm:"index"`
	PaymentPointer string    `json:"paymentPointer"`
	RequestID      string    `json:"requestId"`
	Amount         string    `json:"amount"`
	AssetCode      string    `json:"assetCode"`
	AssetScale     int       `json:"assetScale"`
	Receipt        string    `json:"receipt"`
	AmountValue    float64   `json:"amountValue"`
}

type Payments struct {
	Details []PaymentDetail `json:"payments"`
}

func (d *PaymentDetail) Transform() {
	amount, err := strconv.ParseFloat(d.Amount, 32)
	scale := math.Pow10(d.AssetScale)

	if err != nil {
		log.Fatal(err)
		return
	}

	d.AmountValue = (amount / scale)
}

func (p *Payments) TransformAll () {
	for _, v := range p.Details {
		v.Transform()
	}
}
