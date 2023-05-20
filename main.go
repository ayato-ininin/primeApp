package main

import "fmt"

func main() {
	n := 0
	_, msg := isPrime(n)
	fmt.Println(msg)
}

// 素数かどうかを判定する関数
func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition", n)
	}

	// negative numbers are not prime by definition
	if n < 0 {
		return false, fmt.Sprintf("%d is not prime, because of negative number", n)
	}

	// use the modules operator repeatedly to see if we have a prime numbe
	for i:= 2; i <= n/2; i++ {
		if n % i == 0 {
			return false, fmt.Sprintf("%d is not prime because it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is prime!", n)
}
