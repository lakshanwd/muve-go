package dao

//Fare - holds fare data
type Fare struct {
	FareID        int64          `json:"fareID,omitempty"`
	PaymentOption *PaymentOption `json:"paymentOption,omitempty"`
	BillAmount    float64        `json:"billAmount,omitempty"`
}
