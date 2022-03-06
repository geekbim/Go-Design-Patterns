package main

import "fmt"

type email struct {
	otp
}

func (e *email) getRandomOTP(len int) string {
	randomOtp := "1234"

	fmt.Printf("Email: generating random otp %s\n", randomOtp)

	return randomOtp
}

func (e *email) saveOTPCache(otp string) {
	fmt.Printf("Email: saving otp: %s to cache\n", otp)
}

func (e *email) getMessage(otp string) string {
	return "Email OTP for login is " + otp
}

func (e *email) sendNotification(message string) error {
	fmt.Printf("Email: sending email: %s\n", message)

	return nil
}

func (e *email) publishMetric() {
	fmt.Printf("Email: publishing metrics\n")
}
