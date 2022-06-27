package esending

import "gopkg.in/gomail.v2"

//UseCase use case interface
type SMTPRepo interface {
	DialAndSendMail(m *gomail.Message) (erCode int)
}

type ESending interface {
	Send() (erCode int)
}
