package dao

//Driver - holds driver
type Driver struct {
	DriverID   int64  `json:"driverID,omitempty"`
	DriverName string `json:"driverName,omitempty"`
}
