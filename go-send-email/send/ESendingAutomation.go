package send

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/ethanol1310/send-template-emails/go-send-email/send/common"
	"github.com/gocarina/gocsv"
	"gopkg.in/gomail.v2"
)

type ESendingAutomation struct {
	ListMail          map[CustomerInfo]Mail
	TemplatePath      string
	CustomerPath      string
	ErrorCustomerPath string
	Output            string
}

func (e *ESendingAutomation) PrepareCustomer() {
	in, err := os.Open(e.CustomerPath)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	customers := []*CustomerInfo{}
	if err := gocsv.UnmarshalFile(in, &customers); err != nil {
		panic(err)
	}

	e.ListMail = make(map[CustomerInfo]Mail, len(customers))
	templateMail := Mail{}
	templateMail.ReadTemplate(e.TemplatePath)

	var listMail []Mail
	for _, customer := range customers {
		if customer.ValidCustomer() {
			e.ListMail[*customer] = e.PrepareContent(templateMail, *customer)
			listMail = append(listMail, e.ListMail[*customer])

		} else {
			if !common.FileExists(e.ErrorCustomerPath) {
				common.WriteStringToFile(e.ErrorCustomerPath, "TITLE,FIRST_NAME,LAST_NAME,EMAIL\n", false)
			}
			common.WriteStringToFile(e.ErrorCustomerPath, customer.Title+","+customer.FirstName+","+customer.LastName+","+customer.Email+"\n", true)
		}
	}
	common.WriteDataToJson(e.Output, listMail)
}

func (e *ESendingAutomation) PrepareContent(templateMail Mail, customerInfo CustomerInfo) (mail Mail) {
	mail.From = templateMail.From
	mail.MimeType = templateMail.MimeType
	mail.Subject = templateMail.Subject
	mail.GenerateContent(templateMail.Body, customerInfo)
	return mail
}

func (e *ESendingAutomation) Send() {
	fmt.Printf("%d\n", len(e.ListMail))
	for customer, mail := range e.ListMail {
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
