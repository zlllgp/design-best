package slog

import (
	"log"
	"net/smtp"
	"os"
)

// 配置SMTP服务器信息
const (
	SMTPServer   = "smtp.example.com"
	SMTPPort     = "587"
	SMTPUser     = "your-email@example.com"
	SMTPPassword = "your-password"
)

func notifyByEmail(subject, message string) {
	auth := smtp.PlainAuth("", SMTPUser, SMTPPassword, SMTPServer)
	to := []string{"admin@example.com"}
	msg := []byte("To: admin@example.com\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		message + "\r\n")
	err := smtp.SendMail(SMTPServer+":"+SMTPPort, auth, SMTPUser, to, msg)
	if err != nil {
		log.Println("Error sending email:", err)
	}
}

// 错误跟踪和用户通知
func EmailLog() {
	file, err := os.OpenFile("../../log/errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// 示例错误情况
	testError := func() {
		defer func() {
			if r := recover(); r != nil {
				errMsg := "Recovered in testError: " + r.(string)
				logger.Println(errMsg)
				notifyByEmail("Error Notification", errMsg)
			}
		}()
		panic("something went wrong")
	}

	testError()
}
