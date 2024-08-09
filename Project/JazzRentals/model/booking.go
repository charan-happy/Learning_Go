package model

import (
	"time"
)

// DO NOT CHANGE STRUCT SIGNATURE AND DELETE/COMMENT 
type Booking struct {
	BookingID string
	CustomerName string
	TripType  string
	TripDate  time.Time
	Duration     int
	VehicleList  []Vehicle
}
