package esending

import (
	"fmt"
	"os"

	"github.com/ethanol1310/send-template-emails/go-send-email/internal/entity"
	"github.com/ethanol1310/send-template-emails/go-send-email/pkg/common"
	"github.com/ethanol1310/send-template-emails/go-send-email/pkg/helper"
	"github.com/gocarina/gocsv"
	"gopkg.in/gomail.v2"
)

type ESendingTemplate struct {
	Dialer   SMTPRepo
	ListMail map[entity.CustomerInfo]*entity.Mail
}

func New(d SMTPRepo) *ESendingTemplate {
	return &ESendingTemplate{
		Dialer: d,
	}
}

func (s *ESendingTemplate) PrepareMailToSend(e entity.ESendingAutomation) (erCode int) {
	in, err := os.Open(e.CustomerPath)
	if err != nil {
		return common.MKFAIL(common.NOT_OPEN)
	}
	defer in.Close()

	customers := []*entity.CustomerInfo{}
	if err := gocsv.UnmarshalFile(in, &customers); err != nil {
		return common.MKFAIL(common.NOT_MARSHALL)
	}

	s.ListMail = make(map[entity.CustomerInfo]*entity.Mail, len(customers))
	templateMail := entity.Mail{}
	erCode = templateMail.ReadTemplateMail(e.TemplatePath)
	if !common.IS_SUCCESS(erCode) {
		return erCode
	}

	var listMail []*entity.Mail
	for _, customer := range customers {
		if customer.ValidCustomer() {
			mail := &entity.Mail{}
			s.ListMail[*customer], erCode = mail.GenerateMailFromTemplate(templateMail, *customer)
			if common.IS_SUCCESS(erCode) {
				listMail = append(listMail, s.ListMail[*customer])
			} else {
				// ::TODO
				// Log
			}
		} else {
			if !helper.FileExists(e.ErrorCustomerPath) {
				erCode = helper.WriteStringToFile(e.ErrorCustomerPath, "TITLE,FIRST_NAME,LAST_NAME,EMAIL\n", false)
				if !common.IS_SUCCESS(erCode) {
					return erCode
				}
			}
			erCode = helper.WriteStringToFile(e.ErrorCustomerPath, customer.Title+","+customer.FirstName+","+customer.LastName+","+customer.Email+"\n", true)
			if !common.IS_SUCCESS(erCode) {
				return erCode
			}
		}
	}
	erCode = helper.WriteDataToJson(e.Output, listMail)
	if !common.IS_SUCCESS(erCode) {
		return erCode
	}
	return common.MKSUCCESS()
}

func (s *ESendingTemplate) Send() (erCode int) {
	for customer, mail := range s.ListMail {
		m := gomail.NewMessage()

		m.SetHeader("To", customer.Email)
		m.SetHeader("Subject", mail.Subject)
		m.SetBody("text/plain", mail.Body)

		er := s.Dialer.DialAndSendMail(m)
		fmt.Printf("Send - erCode=%d: %s\n", er, customer.Email)
	}

	return common.MKSUCCESS()
}
