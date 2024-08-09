package main

import (
	"JazzRentals/model"
	"JazzRentals/service"
	"fmt"
	"strconv"
	"time"
)

func main() {
	//Uncomment the below functions after implementing the project
	//bookTrip()
	//getVehicleReport()
}

// DO NOT CHANGE METHOD SIGNATURE AND DELETE/COMMENT METHOD
func BookTrip() {

	car1 := model.Car{
		ModelName:   "Creta",
		CarType: "SUV",
	}
	bike1 := model.Bike{
		ModelName:   "Platina",
		Color: "Blue",
	}
	booking1 := model.Booking{
		CustomerName: "Alice",
		TripType:  "Out-Station",
		TripDate:  time.Now().AddDate(0, 0, 2),
		Duration:     5,
		VehicleList:  []model.Vehicle{car1, bike1},
	}
	bookingPrice, bookingID, err := service.BookTrip(&booking1)
	if bookingID!="NA" && err==nil {
		fmt.Println("Your trip has been booked successfully!!")
		fmt.Println("Booking ID : ", bookingID, "\nBooking Price : ", bookingPrice)
	} else {
		fmt.Println("\n\tERROR:", err)
	}
}

// DO NOT CHANGE METHOD SIGNATURE AND DELETE/COMMENT METHOD
func GetVehicleReport() {
	vehicleReport, err := service.GetVehicleReport("Dezire")
	if len(vehicleReport) != 0 && err == nil {
		fmt.Println("\n\t...Fetching Vehicle Report....")
		fmt.Println("\n\tModelName       BookingStatus")
		fmt.Println("--------------------------------------------------")
		for _, report := range vehicleReport {

			fmt.Println("\t " + report.ModelName +  "\t\t" + strconv.FormatBool(report.BookingStatus) + " \t\t")
		}
	} else {
		fmt.Println("\n\tERROR:", err)
		fmt.Println()
	}

}
