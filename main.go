package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print a welcome message
	intro()

	// create a channel to indicat e when the program can quit
	doneChan := make(chan bool)

	// start a gorouutin to read user input and run program
	go readUserInput(doneChan)

	// block until the done chan gets a value
	<-doneChan

	// close the channel
	close(doneChan)

	// say goodbye
	fmt.Println("Goodbye!")
}

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// try to convert what the user typed into an int
	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole numnber!", false
	}

	_, msg := isPrime(numToCheck)
	return msg, false
}

func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("-----------")
	fmt.Println("Enter a whole number, and we'll tell you if it's prime or not.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
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
