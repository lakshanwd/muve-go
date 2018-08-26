package dao

//PaymentOption - holds payemnt option data
type PaymentOption struct {
	PaymentOptionID int64      `json:"paymentOptionID,omitempty"`
	Passenger       *Passenger `json:"passenger,omitempty"`
	PaymentType     string     `json:"paymentType,omitempty"`
}
