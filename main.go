package main

import (
	"fmt"
	"regexp"
	"strings"
)

func check(c_card string) bool {
	digits := strings.ReplaceAll(c_card, "-", "")
	for i := 0; i < len(digits)-3; i++ {
		if digits[i] == digits[i+1] && digits[i] == digits[i+2] && digits[i] == digits[i+3] {
			return false
		}
	}
	return true
}

func main() {
	var pattern1 = regexp.MustCompile(`^[4-6][0-9]{15}$`)
	var pattern2 = regexp.MustCompile(`^[4-6][0-9]{3}(-[0-9]{4}){3}$`)

	var number_creditcards int
	fmt.Scan(&number_creditcards)

	results := make([]string, number_creditcards)
	var credit_card string

	for i := 0; i < number_creditcards; i++ {
		fmt.Scan(&credit_card)
		if (pattern1.MatchString(credit_card) || pattern2.MatchString(credit_card)) && check(credit_card) {
			results[i] = "Valid"
		} else {
			results[i] = "Invalid"
		}
	}
	for _, result := range results {
		fmt.Println(result)
	}
}
