package card_validator_test

import (
	"testing"

	card_validator "github.com/JrSchmidtt/card-validator"
	"github.com/stretchr/testify/assert"
)

var creditCardNumbers = map[string]map[string]string{
	"master": {
		"valid":   "5158-8200-1179-0888",
		"invalid": "1234-1234-1234-1234",
	},
	"visa": {
		"valid":   "9876-5432-1098-7654",
		"invalid": "9876-5432-1098-7652",
	},
	"amex": {
		"valid":   "1111-2222-3333-4444",
		"invalid": "1111-2222-3333-4442",
	},
	"discover": {
		"valid":   "5555-6666-7777-8888",
		"invalid": "5555-6666-7777-8882",
	},
	"diners": {
		"valid":   "9999-0000-1111-2222",
		"invalid": "9999-0000-1111-2223",
	},
}

type testcase struct {
	scenario string
	function func(t *testing.T)
}

var validator card_validator.CardValidator

func TestMain(m *testing.M) {
	validator = card_validator.NewCardValidator()
	m.Run()
}

func TestCardValidator(t *testing.T) {
	testcases := []testcase{
		{
			scenario: "Given a valid card number, the Validate method should return true",
			function: testSuccess,
		},
		{
			scenario: "Given an invalid card number, the Validate method should return false",
			function: testInvalidCardNumber,
		},
		{
			scenario: "Card number is too short",
			function: testShortCardNumber,
		},
		{
			scenario: "Card number is too long",
			function: testLongCardNumber,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.scenario, tc.function)
	}
}

const (
	errExpectedCardNumberToBeValid   = "Expected card number to be valid"
	errExpectedCardNumberToBeInvalid = "Expected card number to be invalid"
	errExpectedBrandNameToBeEmpty    = "Expected brand name to be empty"
)

func testSuccess(t *testing.T) {
	isValid, brandName := validator.Validate(creditCardNumbers["master"]["valid"])
	assert.True(t, brandName == "master", "Expected brand name to be master")
	assert.True(t, isValid, errExpectedCardNumberToBeValid)
}

func testShortCardNumber(t *testing.T) {
	isValid, brandName := validator.Validate("1234")
	assert.False(t, isValid, errExpectedCardNumberToBeInvalid)
	assert.True(t, brandName == "", errExpectedBrandNameToBeEmpty)
}

func testLongCardNumber(t *testing.T) {
	isValid, brandName := validator.Validate("12345678901234567890")
	assert.False(t, isValid, errExpectedCardNumberToBeInvalid)
	assert.True(t, brandName == "", errExpectedBrandNameToBeEmpty)
}

func testInvalidCardNumber(t *testing.T) {
	isValid, brandName := validator.Validate(creditCardNumbers["master"]["invalid"])
	assert.False(t, isValid, errExpectedCardNumberToBeInvalid)
	assert.True(t, brandName == "", errExpectedBrandNameToBeEmpty)
}
