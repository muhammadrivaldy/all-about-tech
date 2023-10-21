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
	otp := fmt.Sprint(goutil.RandIntRunes(6))

	otpRepository := OTPRepository{}
	otpRepository.SaveOTP(phoneNumber, otp)

	notificationExternal := NotificationExternal{}
	notificationExternal.SendToWhatsApp(phoneNumber, otp)
}

type OTPRepository struct{}

func (o *OTPRepository) SaveOTP(phoneNumber string, otp string) {
	/* {Code for saving the data into database} */
	fmt.Printf("This is the phone number %s and the otp %s that we will save into the database\n", phoneNumber, otp)
}

type NotificationExternal struct{}

func (n *NotificationExternal) SendToWhatsApp(phoneNumber string, otp string) {
	/* {Code for sending the otp to the user's phone number} */
	fmt.Printf("Sending otp %s via whatsapp to the phone number %s", otp, phoneNumber)
}

func main() {
	otpController := OTPController{}
	otpController.SendOTPEndpoint()
}
