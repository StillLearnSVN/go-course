package main


func main() {
	// Switch case example
	day := "Tuesday"
	weather := "Sunny"

	switch {
	case day == "Monday" && weather == "Rainy":
		println("Start of the work week with a gloomy weather!")
	case day == "Monday" && weather == "Sunny":
		println("Start of the work week with a bright day!")
	case day == "Tuesday" && weather == "Sunny":
		println("Second day of the work week and the sun is shining!")
	case day == "Tuesday" && weather == "Cloudy":
		println("Second day of the work week with some clouds.")
	case day == "Friday" && weather == "Sunny":
		println("Almost the weekend and the weather is perfect!")
	case day == "Friday" && weather == "Rainy":
		println("Almost the weekend, but it's raining.")
	default:
		println("It's just another day with unpredictable weather.")
	}
}