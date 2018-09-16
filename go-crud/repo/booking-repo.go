package repo

import (
	"container/list"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/lakshanwd/muve-go/go-crud/dao"
)

//GetBookings - get bookings with pagination
func GetBookings(from int64, count int) (*list.List, error) {
	query := "select tb.booking_id, tb.pick_up_address, tb.drop_address, tb.total_distance, tb.created_on, tdv.vehicle_no, td.driver_name, tp.name, tp.phone, if (isnull(tf.bill_amount),0,tf.bill_amount) as bill_amount from tbl_booking as tb INNER JOIN tbl_driver_vehicle as tdv on tdv.vehicle_id=tb.vehicle_id INNER JOIN tbl_driver as td on td.driver_id=tdv.driver_id INNER JOIN tbl_passenger as tp on tp.passenger_id=tb.passenger_id LEFT JOIN tbl_fare as tf ON tf.fare_id=tb.fare_id WHERE status=1 order by tb.booking_id limit ?,?"
	readerFn := func(rows *sql.Rows, collection *list.List) {
		var booking dao.Booking
		var vehicle dao.DriverVehicle
		var driver dao.Driver
		var fare dao.Fare
		var passenger dao.Passenger
		err := rows.Scan(&booking.BookingID, &booking.PickUpAddress, &booking.DropAddress, &booking.TotalDistance, &booking.CreatedOn, &vehicle.VehicleNo, &driver.DriverName, &passenger.Name, &passenger.Phone, &fare.BillAmount)
		if err == nil {
			vehicle.Driver = &driver
			booking.Fare = &fare
			booking.DriverVehicle = &vehicle
			booking.Fare = &fare
			booking.Passenger = &passenger
			collection.PushBack(booking)
		} else {
			log.Fatal(err)
		}
	}
	return Select(query, readerFn, from, count)
}

//InsertBooking - insert booking
func InsertBooking(booking *dao.Booking) (int64, error) {
	query := "INSERT INTO tbl_booking(vehicle_id, pick_up_address, drop_address, total_distance, created_on, passenger_id, fare_id) values (?,?,?,?,?,?,?)"
	t := time.Now()
	timeString := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fareID := func() interface{} {
		if booking.Fare != nil {
			return booking.Fare.FareID
		}
		return nil
	}
	return Insert(query, &booking.DriverVehicle.VehicleID, &booking.PickUpAddress, &booking.DropAddress, &booking.TotalDistance, timeString, &booking.Passenger.PassengerID, fareID())
}

//GetBooking - get single booking
func GetBooking(id int64) (*list.List, error) {
	query := "select tb.booking_id, tb.pick_up_address, tb.drop_address, tb.total_distance, tb.created_on, tdv.vehicle_no, td.driver_name, tp.name, tp.phone, if(isnull(tf.bill_amount),0,tf.bill_amount) as bill_amount from tbl_booking as tb INNER JOIN tbl_driver_vehicle as tdv on tdv.vehicle_id=tb.vehicle_id INNER JOIN tbl_driver as td on td.driver_id=tdv.driver_id INNER JOIN tbl_passenger as tp on tp.passenger_id=tb.passenger_id LEFT JOIN tbl_fare as tf ON tf.fare_id=tb.fare_id where tb.booking_id = ? and status=1"
	readerFn := func(rows *sql.Rows, collection *list.List) {
		var booking dao.Booking
		var vehicle dao.DriverVehicle
		var driver dao.Driver
		var fare dao.Fare
		var passenger dao.Passenger
		err := rows.Scan(&booking.BookingID, &booking.PickUpAddress, &booking.DropAddress, &booking.TotalDistance, &booking.CreatedOn, &vehicle.VehicleNo, &driver.DriverName, &passenger.Name, &passenger.Phone, &fare.BillAmount)
		if err == nil {
			vehicle.Driver = &driver
			booking.Fare = &fare
			booking.DriverVehicle = &vehicle
			booking.Fare = &fare
			booking.Passenger = &passenger
			collection.PushBack(booking)
		} else {
			log.Fatal(err)
		}
	}
	return Select(query, readerFn, id)
}

//UpdateBooking - update existing booking
func UpdateBooking(booking *dao.Booking, id int) (int64, error) {
	query := "Update tbl_booking set vehicle_id=?, pick_up_address=?, drop_address=?, total_distance=?, created_on=?, passenger_id=?, fare_id=? WHERE booking_id=? and status=1"
	fareID := func() interface{} {
		if booking.Fare != nil {
			return booking.Fare.FareID
		}
		return nil
	}
	return UpdateDelete(query, &booking.DriverVehicle.VehicleID, &booking.PickUpAddress, &booking.DropAddress, &booking.TotalDistance, &booking.CreatedOn, &booking.Passenger.PassengerID, fareID(), id)
}

//RemoveBooking - remove existing booking
func RemoveBooking(id int) (int64, error) {
	query := "Update tbl_booking set status=0 WHERE booking_id=? and status=1"
	return UpdateDelete(query, id)
}
