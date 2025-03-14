package card_validator

import (
	"regexp"
)

type CardValidator interface {
	Validate(cardNumber string) (isValid bool, brandName string)
}

type cardValidator struct {
	brandPrefixes map[string]string
}

func NewCardValidator(customPrefixes ...map[string]string) CardValidator {
	var defaultPrefixes = map[string]string{
		"4":    "visa",
		"5":    "master",
		"34":   "amex",
		"37":   "amex",
		"6011": "discover",
		"380":  "diners",
		"386":  "diners",
		"388":  "diners",
	}
	if len(customPrefixes) > 0 {
		for prefix, brand := range customPrefixes[0] {
			defaultPrefixes[prefix] = brand
		}
	}
	return &cardValidator{brandPrefixes: defaultPrefixes}
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

// getBrand returns the brand of a credit card number
// based on the first digits of the card number
func (c *cardValidator) getBrand(cardNumber string) string {
	for prefix, brand := range c.brandPrefixes {
		if len(cardNumber) >= len(prefix) && cardNumber[:len(prefix)] == prefix {
			return brand
		}
	}
	return ""
}
