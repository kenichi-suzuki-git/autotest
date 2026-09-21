package main

// EvenOrOdd determines if a number is even or odd
func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
