package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
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

func Test_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set Os.Stdout to our write pipe(標準出力をバッファにリダイレクト)
	// 基本はコンソール、ターミナルの標準出力に送られるため。
	os.Stdout = w

	prompt()

	// close the write pipe
	_ = w.Close()

	// reset os.Stdout to what it was before(バッファへリダイレクトしたままだと後続のテスト等で正常に動かない)
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expeted -> but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set Os.Stdout to our write pipe(標準出力をバッファにリダイレクト)
	// 基本はコンソール、ターミナルの標準出力に送られるため。
	os.Stdout = w

	intro()

	// close the write pipe
	_ = w.Close()

	// reset os.Stdout to what it was before(バッファへリダイレクトしたままだと後続のテスト等で正常に動かない)
	os.Stdout = oldOut

	// read the output of our prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// Split output into lines
	output := strings.Split(strings.TrimSpace(string(out)), "\n")
	expected := []string{
		"Is it Prime?",
		"-----------",
		"Enter a whole number, and we'll tell you if it's prime or not.",
		"->",
	}

	if len(output) != len(expected) {
		t.Errorf("Output line count mismatch: got %v want %v", len(output), len(expected))
	}

	for i, line := range output {
		if line != expected[i] {
			t.Errorf("Unexpected output on line %v: got %v want %v", i+1, line, expected[i])
		}
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "zero", input: "0", expected: "0 is not prime, by definition"},
		{name: "one", input: "1", expected: "1 is not prime, by definition"},
		{name: "two", input: "2", expected: "2 is prime!"},
		{name: "three", input: "3", expected: "3 is prime!"},
		{name: "negative", input: "-11", expected: "-11 is not prime, because of negative number"},
		{name: "typed", input: "three", expected: "Please enter a whole number!"},
		{name: "decimal", input: "3.14", expected: "Please enter a whole number!"},
		{name: "quit", input: "q", expected: ""},
		{name: "QUIT", input: "Q", expected: ""},
		{name: "greek", input: "Ω", expected: "Please enter a whole number!"},
	}

	for _, tt := range tests {
		input := strings.NewReader(tt.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)

		if !strings.EqualFold(res, tt.expected) {
			t.Errorf("%s: expected %s but got %s", tt.name, tt.expected, res)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	// to test this. function, we need a channel, and an instance of an io.Reader(引数)
	doneChan := make(chan bool)

	// create a reference to a bytes.Buffer
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
