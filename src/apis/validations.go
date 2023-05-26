package apis

import "regexp"

func validateOTP(otp string) error {
	// checks if OTP is 6 digit in length
	otpPattern := regexp.MustCompile(`^\d{6}$`)
	if otpPattern.MatchString(otp) {
		return nil
	}

	return errInvalidOTP
}

func validatePhoneNumber(phoneNumber string) error {
	// Validate Indian phone numbers having +91 country code
	phonePattern := regexp.MustCompile(`^\+91\d{10}$`)
	if phonePattern.MatchString(phoneNumber) {
		return nil
	}

	return errInvalidPhoneNum
}
