package main

import (
	"fmt"
)

func fizz_buzz(num int) string {
	switch {
	case num % 3 == 0 && num % 5 == 0:
		return "FizzBuzz!"
	case num % 3 == 0:
		return "Fizz"
	case num % 5 == 0:
		return "Buzz"
	default:
		return string(num)
	}
}

func main() {
	spec()
}

func spec() {
	fmt.Println("num: 1", fizz_buzz(1))	
	fmt.Println("num: 3", fizz_buzz(3))	
	fmt.Println("num: 4", fizz_buzz(4))	
	fmt.Println("num: 5", fizz_buzz(5))	
	fmt.Println("num: 7", fizz_buzz(7))	
	fmt.Println("num: 12", fizz_buzz(12))	
	fmt.Println("num: 15", fizz_buzz(15))	
	fmt.Println("num: 16", fizz_buzz(16))	
	fmt.Println("num: 18", fizz_buzz(18))	
	fmt.Println("num: 19", fizz_buzz(19))	
}