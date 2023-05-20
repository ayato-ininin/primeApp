package main

import "testing"

func Test_isPrime(t *testing.T) {
	// result, msg := isPrime(0)
	// if result {
	// 	t.Errorf("with %d as test parameter, got true, but expected false", 0)
	// }

	// if msg != "0 is not prime, by definition" {
	// 	t.Errorf("with %d as test parameter, got %s, but expected %s", 0, msg, "0 is not prime, by definition")
	// }

	// result, msg = isPrime(7)
	// if !result {
	// 	t.Errorf("with %d as test parameter, got true, but expected false", 7)
	// }

	// if msg != "7 is prime!" {
	// 	t.Errorf("with %d as test parameter, got %s, but expected %s", 7, msg, "7 is prime!")
	// }

	// Table Test
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"zero", 0, false, "0 is not prime, by definition"},
		{"one", 1, false, "1 is not prime, by definition"},
		{"negative number", -11, false, "-11 is not prime, because of negative number"},
		{"prime", 7, true, "7 is prime!"},
		{"not prime", 8, false, "8 is not prime because it is divisible by 2"},
	}

	for _, tt := range primeTests {
		result, msg := isPrime(tt.testNum)
		if tt.expected && !result {
			t.Errorf("%s: expected true but got false", tt.name)
		}
		if !tt.expected && result {
			t.Errorf("%s: expected false but got true", tt.name)
		}
		if tt.msg != msg {
			t.Errorf("%s: expected %s but got %s", tt.name, tt.msg, msg)
		}
	}
}
