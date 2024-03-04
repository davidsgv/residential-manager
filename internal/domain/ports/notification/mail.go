package notification

type MailNotification interface {
	// SendMail(to []string, subject string, body []byte) error
	SendHTMLMail(to []string, subject string, bodyData map[string]any, templateDir ...string) error
}
