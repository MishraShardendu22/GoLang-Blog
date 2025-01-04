package utils

import (
	"fmt"
	"net/smtp"
	"os"
	"sync"
)

func MailSender(sender string, otp int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure that wg.Done() is called regardless of success or failure

	from := os.Getenv("MAIL_ID")
	pass := os.Getenv("MAIL_PASS")
	if from == "" || pass == "" {
		fmt.Println("Environment variables MAIL_ID or MAIL_PASS are not set.")
		os.Exit(1)
	}
	to := sender
	host := "smtp.gmail.com"
	port := "587"
	
	message := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Welcome Email</title>
			<style>
				/* CSS styles as above */
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<h1>Welcome to Our Blog Platform!</h1>
				</div>
				<div class="content">
					<p>Hello there! 👋</p>
					<p>We're thrilled to have you join our community of writers and thinkers. Our platform is designed to give you the freedom to express yourself authentically.</p>
					<div class="otp-box">
						<p>Your One-Time Password (OTP):</p>
						<p class="otp">%d</p>
						<p><strong>⏰ Important:</strong> This OTP will expire in 10 minutes for security purposes.</p>
					</div>
					<p>Get ready to start sharing your thoughts with the world!</p>
				</div>
				<div class="links">
					<a href="https://github.com/MIshraShardendu22" target="_blank">Visit the Creator - GitHub</a>
					<a href="https://www.linkedin.com/in/shardendumishra22/" target="_blank">Visit the Creator - LinkedIn</a>
				</div>
				<div class="footer">
					<p>With creativity and passion,<br>The Blog Team</p>
					<p>If you didn't request this email, please ignore it.</p>
				</div>
			</div>
		</body>
		</html>
	`, otp)

	body := []byte("Subject: Welcome to Our Blog Platform\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" + message)
	auth := smtp.PlainAuth("", from, pass, host)
	err := smtp.SendMail(host+":"+port, auth, from, []string{to}, body)
	if err != nil {
		fmt.Println("Error While Sending Mail:", err)
		os.Exit(1)
	}
	fmt.Println("Mail Sent Successfully")
}

func SendEmailFast(email string, otp int) {
	var wg sync.WaitGroup
	wg.Add(1)
	go MailSender(email, otp, &wg)
	wg.Wait()
}