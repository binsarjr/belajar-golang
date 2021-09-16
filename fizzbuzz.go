package main

import "strconv"

func FizzBuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(n)
	}
}

func main() {
	for i := 1; i <= 20; i++ {
		println(FizzBuzz(i))
	}
}
