package main

import (
	"fmt"
	"time"
)

func main() {
	// Current time
	fmt.Println(time.Now())

	// Specific time: August 19, 2019 at 19:19:19 UTC
	specificTime := time.Date(2019, time.August, 19, 19, 19, 19, 0, time.UTC)
	fmt.Println(specificTime)

	// Parse time
	parsedTime, err := time.Parse("2006-01-02 15:04:05", "2019-08-19 19:19:19")
	if err != nil {
		fmt.Println("Error parsing time:", err)
	} else {
		fmt.Println(parsedTime)
	}

	parsedTime2, _ := time.Parse("06-01-02", "19-08-19")
	fmt.Println(parsedTime2)

	parsedTime3, _ := time.Parse("2006-Jan-02", "2019-Aug-19")
	fmt.Println(parsedTime3)

	// Format time
	t := time.Now()
	formattedTime := t.Format("2006-01-02 15:04:05")
	fmt.Println(formattedTime)

	formattedTime2 := t.Format("06-01-02")
	fmt.Println(formattedTime2)

	oneDayLater := t.Add(24 * time.Hour)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())

	fmt.Println("Rounded Time:", t.Round(24*time.Hour))


	loc, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println("Location:", loc)
	fmt.Println("Time in Shanghai:", t.In(loc))

	// If you want to see the difference in time zones
	nyLoc, _ := time.LoadLocation("America/New_York")
	fmt.Println("Time in New York:", t.In(nyLoc))

	// If you want to see other time zones list
	// do cd$(go env GOROOT)/lib/time/zoneinfo.zip
	// and unzip it, then you can see all the time zones

	// Time subtraction
	t1 := time.Date(2021, time.January, 1, 10, 10, 10, 5, time.UTC)
	t2 := time.Date(2021, time.January, 1, 14, 40, 20, 10, time.UTC)
	diff := t2.Sub(t1)
	fmt.Println("Difference:", diff)
	fmt.Println("Hours:", diff.Hours())
	fmt.Println("Minutes:", diff.Minutes())
	fmt.Println("Seconds:", diff.Seconds())
	fmt.Println("Nanoseconds:", diff.Nanoseconds())

	// Sleep for 2 seconds
	fmt.Println("Sleeping for 2 seconds...")
	time.Sleep(2 * time.Second)
	fmt.Println("Awake now!")

	//Compare two times
	fmt.Println("t1 before t2:", t1.Before(t2))
	fmt.Println("t1 after t2:", t1.After(t2))
	fmt.Println("t1 equal t2:", t1.Equal(t2))

	// Conclusion
	// The time package in Go provides a comprehensive set of functions to handle time-related operations, including getting the current time, formatting and 
	// parsing time strings, manipulating time values, handling time zones, and performing arithmetic operations on time values.
}
