package repo

import (
	"fmt"

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
	if err := sp.DialAndSend(m); err != nil {
		fmt.Println(err)
		return common.MKFAIL(common.FAILED)
	}
	return common.MKSUCCESS()
}
