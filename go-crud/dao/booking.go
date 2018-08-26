package dao

//Booking - holds booking data
type Booking struct {
	BookingID     int64          `json:"bookingID,omitempty"`
	PickUpAddress string         `json:"pickUpAddress,omitempty"`
	DropAddress   string         `json:"dropAddress,omitempty"`
	TotalDistance float64        `json:"totalDistance,omitempty"`
	CreatedOn     string         `json:"createdOn,omitempty"`
	DriverVehicle *DriverVehicle `json:"driverVehicle,omitempty"`
	Passenger     *Passenger     `json:"passenger,omitempty"`
	Fare          *Fare          `json:"fare,omitempty"`
}
