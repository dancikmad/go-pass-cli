package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	passLength   int
	useSymbols   bool
	useNumbers   bool
	useUppercase bool

	symbols   = "!@#$%^&*()_+{}|:<>?~"
	numbers   = "0123456789"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
)

func main() {
	flag.IntVar(&passLength, "len", 8, "Password length")
	flag.BoolVar(&useSymbols, "symbols", false, "Use symbols")
	flag.BoolVar(&useNumbers, "numbers", false, "Use numbers")
	flag.BoolVar(&useUppercase, "uppercase", false, "Use uppercase")
	flag.Parse()

	fmt.Println("Generating password...")
	fmt.Println(
		"Password generated:",
		generatePassword(passLength, useSymbols, useNumbers, useUppercase),
	)
}

func generatePassword(length int, useSymbols bool, useNumbers bool, useUppercase bool) string {
	rand.Seed(time.Now().UnixNano()) // Ensure different random values on each run

	if length < 1 {
		return ""
	}

	charSet := lowercase
	if useSymbols {
		charSet += symbols
	}
	if useNumbers {
		charSet += numbers
	}
	if useUppercase {
		charSet += uppercase
	}

	// Generate random password
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = charSet[rand.Intn(len(charSet))]
	}

	// Ensure at least one of each required type
	ensureCharacter(&password, useSymbols, symbols)
	ensureCharacter(&password, useNumbers, numbers)
	ensureCharacter(&password, useUppercase, uppercase)

	return string(password)
}

// Ensures at least one character of a specific type is in the password
func ensureCharacter(password *[]byte, condition bool, charSet string) {
	if condition {
		index := rand.Intn(len(*password)) // Pick random position
		(*password)[index] = charSet[rand.Intn(len(charSet))]
	}
}
