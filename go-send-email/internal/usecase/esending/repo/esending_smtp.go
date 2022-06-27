package repo

import (
	"github.com/ethanol1310/send-template-emails/go-send-email/pkg/common"
	"gopkg.in/gomail.v2"
)

type SMTPRepo struct {
	*gomail.Dialer
}

// New -.
func New(sp *gomail.Dialer) *SMTPRepo {
	return &SMTPRepo{sp}
}

func (sp *SMTPRepo) DialAndSendMail(m *gomail.Message) (erCode int) {
	m.SetHeader("From", sp.Username)

	if err := sp.DialAndSend(m); err != nil {
		return common.MKFAIL(common.FAILED)
	}
	return common.MKSUCCESS()
}
