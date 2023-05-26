package apis

import "testing"

func Test_validateOTP(t *testing.T) {
	tests := []struct {
		otp      string
		expected error
	}{
		{otp: "", expected: errInvalidOTP},
		{otp: "123", expected: errInvalidOTP},
		{otp: "123456", expected: nil},
		{otp: "1234567", expected: errInvalidOTP},
	}

	for _, test := range tests {
		result := validateOTP(test.otp)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %s, but got: %s", test.otp, test.expected, result)
		}
	}
}

func Test_validatePhoneNumber(t *testing.T) {
	tests := []struct {
		PhoneNumber string
		expected    error
	}{
		{PhoneNumber: "", expected: errInvalidPhoneNum},
		{PhoneNumber: "+211234567891", expected: errInvalidPhoneNum},
		{PhoneNumber: "+91654", expected: errInvalidPhoneNum},
		{PhoneNumber: "+911234567891", expected: nil},
		{PhoneNumber: "+91234234342234", expected: errInvalidPhoneNum},
	}

	for _, test := range tests {
		result := validatePhoneNumber(test.PhoneNumber)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %s, but got: %s", test.PhoneNumber, test.expected, result)
		}
	}
}
