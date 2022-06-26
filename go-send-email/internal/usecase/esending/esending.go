package esending

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/ethanol1310/send-template-emails/go-send-email/internal/entity"
	"github.com/ethanol1310/send-template-emails/go-send-email/pkg/common"
	"github.com/gocarina/gocsv"
	"gopkg.in/gomail.v2"
)

type Service struct {
	ListMail map[entity.CustomerInfo]*entity.Mail
}

func (s *Service) prepareMailToSend(e entity.ESendingAutomation) {
	in, err := os.Open(e.CustomerPath)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	customers := []*entity.CustomerInfo{}
	if err := gocsv.UnmarshalFile(in, &customers); err != nil {
		panic(err)
	}

	s.ListMail = make(map[entity.CustomerInfo]*entity.Mail, len(customers))
	templateMail := entity.Mail{}
	templateMail.ReadTemplate(e.TemplatePath)

	var listMail []*entity.Mail
	for _, customer := range customers {
		if customer.ValidCustomer() {
			s.ListMail[*customer] = s.prepareMailFromTemplate(templateMail, *customer)
			listMail = append(listMail, s.ListMail[*customer])

		} else {
			if !common.FileExists(e.ErrorCustomerPath) {
				common.WriteStringToFile(e.ErrorCustomerPath, "TITLE,FIRST_NAME,LAST_NAME,EMAIL\n", false)
			}
			common.WriteStringToFile(e.ErrorCustomerPath, customer.Title+","+customer.FirstName+","+customer.LastName+","+customer.Email+"\n", true)
		}
	}
	common.WriteDataToJson(e.Output, listMail)
}

func (s *Service) prepareMailFromTemplate(templateMail entity.Mail, customerInfo entity.CustomerInfo) *entity.Mail {
	mail := entity.Mail{}
	mail.From = templateMail.From
	mail.MimeType = templateMail.MimeType
	mail.Subject = templateMail.Subject
	mail.GenerateMailFromTemplate(templateMail.Body, customerInfo)
	return &mail
}

func (s *Service) Send(e entity.ESendingAutomation) {
	s.prepareMailToSend(e)
	for customer, mail := range s.ListMail {
		fmt.Printf("%s - %s - %s - %s\n", customer.Title, customer.FirstName, customer.LastName, customer.Email)
		m := gomail.NewMessage()

		m.SetHeader("From", common.FROM)
		m.SetHeader("To", customer.Email)
		m.SetHeader("Subject", mail.Subject)
		m.SetBody("text/plain", mail.Body)

		// Settings for SMTP server
		d := gomail.NewDialer("smtp.gmail.com", 587, common.FROM, common.PASSWORD)

		// This is only needed when SSL/TLS certificate is not valid on server.
		// In production this should be set to false.
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		if err := d.DialAndSend(m); err != nil {
			fmt.Println(err)
			panic(err)
		}
		fmt.Printf("SUCCESS\n")
	}
}
