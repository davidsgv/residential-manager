package notification

import (
	"bytes"
	"text/template"

	"gopkg.in/gomail.v2"
)

type EmailConfig struct {
	Mail     string //The email to send mails
	Password string //The password of the mail
	SmptHost string //Host ip
	SmptPort int    //Host port
}

type Mail struct {
	config EmailConfig
}

func NewMail(config EmailConfig) *Mail {
	return &Mail{
		config: config,
	}
}

// func (m *Mail) SendMail(to []string, subject string, body []byte) error {
// 	mail := gomail.NewMessage()
// 	mail.SetHeader("From", m.config.Mail)
// 	mail.SetHeader("To", to...)
// 	mail.SetHeader("Subject", subject)

// 	mail.SetAddressHeader("Cc", "dan@example.com", "Dan")
// 	mail.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
// 	//mail.Attach("/home/Alex/lolcat.jpg")

// 	d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")

// 	// Send the email to Bob, Cora and Dan.
// 	if err := d.DialAndSend(m); err != nil {
// 		panic(err)
// 	}
// }

func (m *Mail) SendHTMLMail(to []string, subject string, bodyData map[string]any, templateDir ...string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", m.config.Mail)
	mail.SetHeader("To", to...)
	mail.SetHeader("Subject", subject)
	//mail.SetAddressHeader("Cc", "dan@example.com", "Dan")

	t, err := template.ParseFiles(templateDir...)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	t.Execute(&body, bodyData)

	mail.SetBody("text/html", body.String())
	//mail.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(m.config.SmptHost, m.config.SmptPort, m.config.Mail, m.config.Password)

	// Send the email to Bob, Cora and Dan.
	err = d.DialAndSend(mail)
	if err != nil {
		return err
	}

	return nil
}
