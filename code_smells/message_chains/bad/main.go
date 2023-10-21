package main

import (
	"fmt"

	goutil "github.com/muhammadrivaldy/go-util"
)

type OTPController struct{}

func (o *OTPController) SendOTPEndpoint() {
	sendOTPService := SendOTPService{}
	sendOTPService.SendOTPViaWhatsApp("+6287728391821")
}

type SendOTPService struct{}

func (s *SendOTPService) SendOTPViaWhatsApp(phoneNumber string) {
	otpService := OTPService{}
	otpService.SendOTP(phoneNumber, fmt.Sprint(goutil.RandIntRunes(6)))
}

type OTPService struct{}

func (o *OTPService) SendOTP(phoneNumber string, otp string) {
	notificationService := NotificationService{}
	notificationService.SendToWhatsApp(phoneNumber, otp)
}

type NotificationService struct{}

func (n *NotificationService) SendToWhatsApp(phoneNumber string, otp string) {
	otpRepository := OTPRepository{}
	otpRepository.SaveOTP(phoneNumber, otp)
}

type OTPRepository struct{}

func (o *OTPRepository) SaveOTP(phoneNumber string, otp string) {
	/* {Code for saving the data into database} */
	fmt.Printf("This is the phone number %s and the otp %s that we will save into the database\n", phoneNumber, otp)
}

func main() {
	otpController := OTPController{}
	otpController.SendOTPEndpoint()
}
