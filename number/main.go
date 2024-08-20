package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var name string

	new := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name:")
	name, _ = new.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Hello %v!\n", name)

	reader := bufio.NewReader(os.Stdin)
	var mobileNumber string

	for {
		fmt.Println("Enter your 10-digit mobile number:")
		input, _ := reader.ReadString('\n')

		// Trim any leading or trailing whitespace
		input = strings.TrimSpace(input)

		// Check if input length is exactly 10
		if len(input) == 10 && isNumeric(input) {
			mobileNumber = input
			break
		} else {
			fmt.Println("Invalid input. Please enter exactly 10 digits.")
		}

	}

	fmt.Printf("Your number is %s\n", mobileNumber)
}

// isNumeric checks if a string contains only numeric characters
func isNumeric(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true

}
