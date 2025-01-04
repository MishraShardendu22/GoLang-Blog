package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func MailSender(sender string, otp int) {
	fmt.Println("Sending Mail...")
	from := os.Getenv("MAIL_ID")
	pass := os.Getenv("MAIL_PASS")
	if from == "" || pass == "" {
		fmt.Println("Environment variables MAIL_ID or MAIL_PASS are not set.")
		os.Exit(1)
	}
	to := sender
	host := "smtp.gmail.com"
	port := "587"
	// Enhanced HTML Email Template
	message := fmt.Sprintf(`
		<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Welcome Email</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f9f9f9;
        }
        .container {
            max-width: 600px;
            margin: 30px auto;
            padding: 20px;
            background: #ffffff;
            border-radius: 12px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
            border: 1px solid #e4e4e4;
        }
        .header {
            text-align: center;
            padding: 20px 0;
            background: linear-gradient(45deg, #6a11cb, #2575fc);
            color: #fff;
            border-radius: 12px 12px 0 0;
        }
        .header h1 {
            margin: 0;
            font-size: 28px;
            font-weight: bold;
        }
        .content {
            padding: 20px;
            line-height: 1.8;
            color: #333;
        }
        .content p {
            margin: 15px 0;
        }
        .otp-box {
            background: #f0f4ff;
            padding: 15px;
            margin: 20px 0;
            text-align: center;
            border-radius: 8px;
            border: 1px dashed #2575fc;
        }
        .otp-box .otp {
            font-size: 24px;
            font-weight: bold;
            color: #2575fc;
        }
        .links {
            margin-top: 20px;
            text-align: center;
        }
        .links a {
            display: inline-block;
            margin: 10px;
            color: #2575fc;
            text-decoration: none;
            font-weight: bold;
            font-size: 14px;
            padding: 10px 15px;
            border: 1px solid #2575fc;
            border-radius: 6px;
            transition: all 0.3s;
        }
        .links a:hover {
            background-color: #2575fc;
            color: #fff;
        }
        .footer {
            margin-top: 30px;
            text-align: center;
            font-size: 14px;
            color: #666;
            padding-top: 20px;
            border-top: 1px solid #e4e4e4;
        }
        .footer p {
            margin: 5px 0;
        }
        .footer a {
            color: #2575fc;
            text-decoration: none;
            font-weight: bold;
        }
        .footer a:hover {
            text-decoration: underline;
        }
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