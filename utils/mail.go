package utils

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

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
	return sendMail(m)
}



func SendApplicationAcceptedMail(toEmail, seekerName, jobTitle, businessName string) error {
    m := gomail.NewMessage()

    m.SetHeader("From", os.Getenv("EMAIL_FROM"))
    m.SetHeader("To", toEmail)
    m.SetHeader("Subject", "Congratulations! Your Application was Accepted - HoodHire")

    body := fmt.Sprintf(`
        <!DOCTYPE html>
        <html>
        <head>
            <style>
                body { font-family: Arial, sans-serif; }
                .container { max-width: 600px; margin: 0 auto; padding: 20px; }
                .banner {
                    background-color: #10b981;
                    padding: 20px;
                    text-align: center;
                    border-radius: 8px;
                    margin: 20px 0;
                }
                .banner h2 { color: white; margin: 0; font-size: 24px; }
                .details {
                    background-color: #f4f4f4;
                    padding: 20px;
                    border-radius: 8px;
                    margin: 20px 0;
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
                <h2>Hey %s!</h2>
                <div class="banner">
                    <h2>🎉 Your Application was Accepted!</h2>
                </div>
                <p>Great news! <strong>%s</strong> has accepted your application for the following position:</p>
                <div class="details">
                    <p><strong>Job:</strong> %s</p>
                    <p><strong>Business:</strong> %s</p>
                </div>
                <p>Log in to HoodHire to view the details and get in touch with the employer.</p>
                <div class="footer">
                    <p>This is an automated email, please do not reply.</p>
                    <p>© HoodHire</p>
                </div>
            </div>
        </body>
        </html>
    `, seekerName, businessName, jobTitle, businessName)

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



func SendApplicationRejectedMail(toEmail, seekerName, jobTitle, businessName string) error {
    m := gomail.NewMessage()

    m.SetHeader("From", os.Getenv("EMAIL_FROM"))
    m.SetHeader("To", toEmail)
    m.SetHeader("Subject", "Application Update - HoodHire")

    body := fmt.Sprintf(`
        <!DOCTYPE html>
        <html>
        <head>
            <style>
                body { font-family: Arial, sans-serif; }
                .container { max-width: 600px; margin: 0 auto; padding: 20px; }
                .banner {
                    background-color: #f4f4f4;
                    padding: 20px;
                    text-align: center;
                    border-radius: 8px;
                    margin: 20px 0;
                }
                .details {
                    background-color: #f4f4f4;
                    padding: 20px;
                    border-radius: 8px;
                    margin: 20px 0;
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
                <h2>Hey %s,</h2>
                <div class="banner">
                    <h2>Application Status Update</h2>
                </div>
                <p>Thank you for your interest. Unfortunately, <strong>%s</strong> has decided to move forward with other candidates for the following position:</p>
                <div class="details">
                    <p><strong>Job:</strong> %s</p>
                    <p><strong>Business:</strong> %s</p>
                </div>
                <p>Don't be discouraged — there are plenty of other opportunities on HoodHire. Keep applying!</p>
                <div class="footer">
                    <p>This is an automated email, please do not reply.</p>
                    <p>© HoodHire</p>
                </div>
            </div>
        </body>
        </html>
    `, seekerName, businessName, jobTitle, businessName)

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




//subscription activated mail 

func SendSubscriptionActivatedMail(toEmail, name, plan string, endDate time.Time) error {
    m := gomail.NewMessage()
    m.SetHeader("From", os.Getenv("EMAIL_FROM"))
    m.SetHeader("To", toEmail)
    m.SetHeader("Subject", "Subscription Activated - HoodHire Pro")

    body := fmt.Sprintf(`
        <!DOCTYPE html>
        <html>
        <body style="font-family: Arial, sans-serif;">
            <div style="max-width: 600px; margin: 0 auto; padding: 20px;">
                <h2>Hey %s!</h2>
                <div style="background-color: #10b981; padding: 20px; text-align: center; border-radius: 8px; margin: 20px 0;">
                    <h2 style="color: white; margin: 0;">🎉 Welcome to HoodHire Pro!</h2>
                </div>
                <p>Your <strong>%s</strong> subscription is now active.</p>
                <div style="background-color: #f4f4f4; padding: 20px; border-radius: 8px;">
                    <p><strong>Plan:</strong> %s</p>
                    <p><strong>Valid until:</strong> %s</p>
                    <p><strong>Benefits:</strong> Unlimited job postings, Verified badge, Priority listing</p>
                </div>
                <p>Start posting unlimited jobs now!</p>
                <div style="margin-top: 30px; font-size: 12px; color: #666;">
                    <p>This is an automated email, please do not reply.</p>
                </div>
            </div>
        </body>
        </html>
    `, name, plan, plan, endDate.Format("Jan 02, 2006"))

    m.SetBody("text/html", body)
    return sendMail(m)
}

//subscription expiry remainder

func SendSubscriptionExpiryReminderMail(toEmail, name, plan string, endDate time.Time) error {
    m := gomail.NewMessage()
    m.SetHeader("From", os.Getenv("EMAIL_FROM"))
    m.SetHeader("To", toEmail)
    m.SetHeader("Subject", "Your HoodHire Pro subscription expires in 7 days")

    body := fmt.Sprintf(`
        <!DOCTYPE html>
        <html>
        <body style="font-family: Arial, sans-serif;">
            <div style="max-width: 600px; margin: 0 auto; padding: 20px;">
                <h2>Hey %s,</h2>
                <div style="background-color: #f59e0b; padding: 20px; text-align: center; border-radius: 8px; margin: 20px 0;">
                    <h2 style="color: white; margin: 0;">⚠️ Subscription Expiring Soon</h2>
                </div>
                <p>Your HoodHire Pro <strong>%s</strong> subscription expires on <strong>%s</strong>.</p>
                <p>Renew now to keep your verified badge and unlimited job postings.</p>
                <div style="margin-top: 30px; font-size: 12px; color: #666;">
                    <p>This is an automated email, please do not reply.</p>
                </div>
            </div>
        </body>
        </html>
    `, name, plan, endDate.Format("Jan 02, 2006"))

    m.SetBody("text/html", body)
    return sendMail(m)
}

//subscription expired mail

func SendSubscriptionExpiredMail(toEmail, name, plan string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", os.Getenv("EMAIL_FROM"))
    m.SetHeader("To", toEmail)
    m.SetHeader("Subject", "Your HoodHire Pro subscription has expired")

    body := fmt.Sprintf(`
        <!DOCTYPE html>
        <html>
        <body style="font-family: Arial, sans-serif;">
            <div style="max-width: 600px; margin: 0 auto; padding: 20px;">
                <h2>Hey %s,</h2>
                <div style="background-color: #ef4444; padding: 20px; text-align: center; border-radius: 8px; margin: 20px 0;">
                    <h2 style="color: white; margin: 0;">Your Pro Subscription Has Expired</h2>
                </div>
                <p>Your HoodHire Pro <strong>%s</strong> subscription has expired.</p>
                <p>Your business verified status has been removed and you are now limited to 5 open jobs.</p>
                <p>Renew your subscription to restore full access.</p>
                <div style="margin-top: 30px; font-size: 12px; color: #666;">
                    <p>This is an automated email, please do not reply.</p>
                </div>
            </div>
        </body>
        </html>
    `, name, plan)

    m.SetBody("text/html", body)
    return sendMail(m)
}

// shared helper to avoid repeating smtp setup
func sendMail(m *gomail.Message) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPassword)

	start := time.Now()

	if err := d.DialAndSend(m); err != nil {
		log.Printf("EMAIL FAILED | to=%v | err=%v | took=%v\n",
			m.GetHeader("To"),
			err,
			time.Since(start),
		)
		return fmt.Errorf("failed to send email: %v", err)
	}

	log.Printf("EMAIL SENT | to=%v | took=%v\n",
		m.GetHeader("To"),
		time.Since(start),
	)

	return nil
}