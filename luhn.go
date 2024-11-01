package luhn

import (
	"fmt"
	"strconv"
)

type luhnNumber struct {
	baseNumber int64
	checkDigit int
}

// calculateCheckDigit calculates the check digit for the given number
func (l luhnNumber) calculateCheckDigit() (luhnNumber, error) {
	if l.baseNumber <= 0 {
		return l, fmt.Errorf("invalid input: number must be positive")
	}

	strNumbers := strconv.FormatInt(l.baseNumber, 10)
	var total int

	for i, multiplier := len(strNumbers)-1, 0; i >= 0; i-- {
		digit, err := strconv.Atoi(string(strNumbers[i]))
		if err != nil {
			return l, fmt.Errorf("invalid digit at position %d: %v", i, err)
		}

		if multiplier%2 == 0 {
			digit *= 2
			total += sumDoubledDigits(digit)
		} else {
			total += digit
		}
		multiplier++
	}

	checkDigit := (10 - (total % 10)) % 10
	l.checkDigit = checkDigit
	return l, nil
}

// FullNumber calculates the full number with the check digit
func FullNumber(number int64) (int64, error) {
	if number <= 0 {
		return 0, fmt.Errorf("invalid input: number must be positive")
	}

	ln := luhnNumber{
		baseNumber: number,
	}
	res, err := ln.calculateCheckDigit()
	if err != nil {
		return 0, fmt.Errorf("failed to generate check digit: %v", err)
	}

	strNumbers := strconv.FormatInt(number, 10) + strconv.Itoa(res.checkDigit)

	finalNumber, err := strconv.ParseInt(strNumbers, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse final number: %v", err)
	}
	return finalNumber, nil
}

// CheckDigit calculates the check digit for the given number
func CheckDigit(number int64) (int, error) {
	if number <= 0 {
		return 0, fmt.Errorf("invalid input: number must be positive")
	}
	ln := luhnNumber{
		baseNumber: number,
	}
	res, err := ln.calculateCheckDigit()
	if err != nil {
		return 0, fmt.Errorf("failed to generate check digit: %v", err)
	}
	return res.checkDigit, nil
}

// IsValid checks if the number is valid according to the Luhn algorithm
func IsValid(number int64) bool {
	s := strconv.FormatInt(number, 10)
	if len(s) < 2 {
		return false
	}

	baseNumberStr := s[:len(s)-1]
	checkDigitStr := s[len(s)-1:]

	baseNumber, _ := strconv.ParseInt(baseNumberStr, 10, 64)
	checkDigit, _ := strconv.Atoi(checkDigitStr)

	ln := luhnNumber{
		baseNumber: baseNumber,
	}
	res, err := ln.calculateCheckDigit()
	if err != nil {
		return false
	}

	return res.checkDigit == checkDigit
}

// sumDoubledDigits sums the digits of a number after doubling
func sumDoubledDigits(digits int) int {
	if digits < 10 {
		return digits
	}

	total := 0
	for digits > 0 {
		total += digits % 10
		digits /= 10
	}
	return total
}
