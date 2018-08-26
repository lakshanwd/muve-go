package dao

//DriverVehicle - holds Driver Vehicle
type DriverVehicle struct {
	VehicleID int64   `json:"vehicleID,omitempty"`
	Driver    *Driver `json:"driver,omitempty"`
	VehicleNo string  `json:"vehicleNo,omitempty"`
}
