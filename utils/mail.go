package utils

import (
	"fmt"
	"os"
	"strconv"
	"math/rand"

	"gopkg.in/gomail.v2"
)


func GenerateOTP() string {
	return fmt.Sprintf("%06d",rand.Intn(1000000))
}

func SendOTPmail(toEmail, otp string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", os.Getenv("EMAIL_FROM"))
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Your OTP Code - HoodHire")

	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<style>
				body { font-family: Arial, sans-serif; }
				.container { max-width: 600px; margin: 0 auto; padding: 20px; }
				.otp-box { 
					background-color: #f4f4f4; 
					padding: 20px; 
					text-align: center; 
					border-radius: 8px;
					margin: 20px 0;
				}
				.otp-code { 
					font-size: 32px; 
					font-weight: bold; 
					color: #333;
					letter-spacing: 5px;
				}
				.footer { 
					margin-top: 30px; 
					font-size: 12px; 
					color: #666; 
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h2>Welcome to HoodHire!</h2>
				<p>Thank you for signing up. Please verify your email address using the OTP below:</p>
				
				<div class="otp-box">
					<div class="otp-code">%s</div>
				</div>
				
				<p>This code will expire in <strong>5 minutes</strong>.</p>
				<p>If you didn't request this code, please ignore this email.</p>
				
				<div class="footer">
					<p>This is an automated email, please do not reply.</p>
				</div>
			</div>
		</body>
		</html>
	`, otp)

	m.SetBody("text/html", body)

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}


