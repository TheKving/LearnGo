package booking

import (
	"fmt"
	"time"
)

func main() {

	//fmt.Println(Schedule("7/25/2019 13:45:00")) // => 2019-07-25 13:45:00 +0000 UTC
	fmt.Println(Schedule("7/25/2019 13:45:00"))
	fmt.Println(HasPassed("July 25, 2019 13:45:00"))
	fmt.Println(IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00"))

	fmt.Println(Description("7/25/2019 13:45:00")) // => "You have an appointment on Thursday, July 25, 2019, at 13:45."
	fmt.Println(AnniversaryDate())                 // => 2020-09-15 00:00:00 +0000 UTC

	fmt.Println(IsAfternoonAppointment("Thursday, May 13, 2010 20:32:00"))
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	dataTime, _ := time.Parse("1/2/2006 15:04:05 ", date) // time.Time, error
	return dataTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	dataTime, _ := time.Parse("January 2, 2006 15:04:05 ", date)
	return time.Now().After(dataTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	dataTime, _ := time.Parse("Monday, January 2, 2006 15:04:05 ", date)
	return dataTime.Hour() >= 12 && dataTime.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	newDate := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", newDate.Weekday(), newDate.Month(), newDate.Day(), newDate.Year(), newDate.Hour(), newDate.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 00, 0, 0, 0, time.UTC)
}
