package main

import (
	"fmt"
	"time"
)

func main() {
	println("Learn Time Package")

	currentTime := time.Now()

	fmt.Println("Current time", currentTime)
	fmt.Printf("Type of currentTime %T\n", currentTime)

	formattedTime := currentTime.Format("02-01-2006") // 02 is day, 01 is month, 2006 is year
	fmt.Println("Formatted Time", formattedTime)

	formattedTime_1 := currentTime.Format("02-01-2006, Monday")
	fmt.Println("Formatted Time with Day Name", formattedTime_1)

	formattedTime_2 := currentTime.Format("02-01-2006, 15:04:05") // 15 is hours, 04 is minutes, 05 is seconds
	fmt.Println("Formatted Time with Time", formattedTime_2)

	formattedTime_3 := currentTime.Format("02-01-2006, 3:04:05") // 3 is 12 hours format
	fmt.Println("Formatted Time with Time", formattedTime_3)

	formattedTime_4 := currentTime.Format("02-01-2006, 3:04:05 PM") // PM AM/PM format
	fmt.Println("Formatted Time with Time", formattedTime_4)

	//Parse Time from string
	dateFormat := "02/01/2006" // DD/MM/YYYY
	dateStr := "25/11/2030"

	formattedTime_5, _ := time.Parse(dateFormat, dateStr)

	fmt.Println("Formated (parse) Date", formattedTime_5) // Formated (parse) Date 2030-11-25 00:00:00 +0000 UTC

	// Always need to same structure like dd/mm/yyyy or dd-mm-yyyy and so on.
	dateFormat_2 := "02, 01, 2006" // DD/MM/YYYY
	dateStr_2 := "25, 11, 2030"

	formattedTime_6, _ := time.Parse(dateFormat_2, dateStr_2)

	fmt.Println("Formated (parse) Date", formattedTime_6) // Formated (parse) Date 2030-11-25 00:00:00 +0000 UTC

	// Add 1 more day in currentDate
	new_date := currentTime.Add(24 * time.Hour) // add duration 24 hours

	fmt.Println("Current time", currentTime)
	fmt.Println("new_date time", new_date)

	formattedTime_7 := currentTime.Format("2006/01/02, Monday")
	fmt.Println("Formatted Time with Day Name", formattedTime_7)
}
