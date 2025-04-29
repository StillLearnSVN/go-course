package basics

func main() {
	// for as while with break
	// sum := 0
	// for {
	// 	sum += 1
	// 	println("Sum: ", sum)
	// 	if sum >= 10 {
	// 		break
	// 	}
	// }
	num := 1
	for num <= 10 {
		if num % 2 == 0 {
			num++
			continue
		}
		println("Odd number: ", num)
		num++
	}
}
