package entity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"time"

	"github.com/ethanol1310/send-template-emails/pkg/common"
)

type Mail struct {
	From     string
	Subject  string
	MimeType string
	Body     string
}

func NewMail(from string, subject string, mime string, body string) *Mail {
	return &Mail{
		From:     from,
		Subject:  subject,
		MimeType: mime,
		Body:     body,
	}
}

func (mail *Mail) generateBodyFromTemplate(templateBody string, customerInfo CustomerInfo) (erCode int) {
	today := time.Now().Format(common.TIME_FORMAT)
	info := TemplateInfo{customerInfo.Title, customerInfo.FirstName, customerInfo.LastName, today}
	mail.Body, erCode = info.ParseString("esending", templateBody)
	return erCode
}

func (mail *Mail) ParseEmailInFrom() {
	// Use regex to parse email from string
	re := regexp.MustCompile(`([a-zA-Z0-9+._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+)`)
	from := re.FindAllString(mail.From, -1)
	if from == nil {
		from = append(from, "default@example.com")
	}
	mail.From = from[0]
}

func (mail *Mail) ReadTemplateMail(filePath string) (errCode int) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("failed to open json file: %s, error: %v", filePath, err)
		return common.MKFAIL(common.NOT_OPEN)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
		return common.MKFAIL(common.NOT_READ)
	}

	if err := json.Unmarshal(jsonData, &mail); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
		return common.MKFAIL(common.NOT_UNMARSHALL)
	}
	return common.MKSUCCESS()
}

func (mail *Mail) GenerateMailFromTemplate(templateMail Mail, customerInfo CustomerInfo) (ret *Mail, erCode int) {
	mail.From = templateMail.From
	mail.MimeType = templateMail.MimeType
	mail.Subject = templateMail.Subject
	erCode = mail.generateBodyFromTemplate(templateMail.Body, customerInfo)
	return mail, erCode
}
