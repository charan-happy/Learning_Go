package repository

import (
	"JazzRentals/model"
)

// DO NOT CHANGE FUNCTION SIGNATURE AND DELETE/COMMENT
func VehicleRepository() []model.VehicleInfo {

	vehiclesInRepository := []model.VehicleInfo{
		{ModelName: "Swift", BookingStatus: true},
		{ModelName: "Dezire", BookingStatus: true},
		{ModelName: "Creta", BookingStatus: true},
	}
	return vehiclesInRepository
}
