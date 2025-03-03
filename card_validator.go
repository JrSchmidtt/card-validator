package card_validator

import (
	"regexp"
)

type CardValidator interface {
	Validate(cardNumber string) (isValid bool, brandName string)
}

type cardValidator struct{}

func NewCardValidator() CardValidator {
	return &cardValidator{}
}

// Validate checks if a credit card number is valid
// A credit card number is valid if it passes the Luhn algorithm
func (c *cardValidator) Validate(rawNumber string) (isValid bool, brandName string) {
	sanitized := c.sanitize(rawNumber)
	if len(sanitized) < 13 || len(sanitized) > 19 {
		return false, ""
	}
	// Luhn algorithm
	var sum int
	isSecond := false
	for i := len(sanitized) - 1; i >= 0; i-- {
		digit := int(sanitized[i] - '0')
		if isSecond {
			digit *= 2
			if digit > 9 {
				digit = digit - 9
			}
		}
		sum += digit
		isSecond = !isSecond
	}
	isValid = sum%10 == 0
	if !isValid {
		return false, ""
	}
	cardBrand := c.getBrand(sanitized)
	return true, cardBrand
}

func (c *cardValidator) sanitize(cardNumber string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(cardNumber, "")
}

func (c *cardValidator) getBrand(cardNumber string) string {
	if cardNumber[0] == '4' {
		return "visa"
	}

	if cardNumber[0] == '5' {
		return "master"
	}

	if cardNumber[0] == '3' && (cardNumber[1] == '4' || cardNumber[1] == '7') {
		return "amex"
	}

	return ""
}
