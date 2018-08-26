package dao

//Passenger - holds passenger data
type Passenger struct {
	PassengerID int64  `json:"passengerID,omitempty"`
	Name        string `json:"name,omitempty"`
	Phone       string `json:"phone,omitempty"`
}
