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
	notificationExternal := NotificationExternal{}
	notificationExternal.SendToWhatsApp(phoneNumber, otp)

	otpRepository := OTPRepository{}
	otpRepository.SaveOTP(phoneNumber, otp)
}

type NotificationExternal struct{}

func (n *NotificationExternal) SendToWhatsApp(phoneNumber string, otp string) {
	/* {Code for sending the otp to the user's phone number} */
	fmt.Printf("Sending otp %s via whatsapp to the phone number %s\n", otp, phoneNumber)
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
