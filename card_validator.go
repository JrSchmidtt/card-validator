package card_validator

import (
	"regexp"
)

type CardValidator interface {
	Validate(cardNumber string) bool
}

type cardValidator struct{}

func NewCardValidator() CardValidator {
	return &cardValidator{}
}

// Validate checks if a credit card number is valid
// A credit card number is valid if it passes the Luhn algorithm
func (c *cardValidator) Validate(rawNumber string) bool {
	sanitized := c.sanitize(rawNumber)

	if len(sanitized) < 13 || len(sanitized) > 19 {
		return false
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
	return sum%10 == 0
}

func (c *cardValidator) sanitize(cardNumber string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(cardNumber, "")
}
