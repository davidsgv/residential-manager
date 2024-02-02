package notification

type MailNotification interface {
	SendMail(from, password, host, port string, to []string) error
}
