package entity_test

import (
	"testing"

	"github.com/ethanol1310/send-template-emails/internal/entity"
	"gopkg.in/go-playground/assert.v1"
)

func TestGenerateMailFromTemplate(t *testing.T) {
	templateMail := entity.Mail{}
	templateMail.ReadTemplateMail("/home/ethanol/Desktop/go/send-template-emails/template/template.json")

	type test struct {
		Title     string
		FirstName string
		LastName  string
		Body      string
		Valid     bool
	}

	tests := []test{
		{
			Title:     "",
			FirstName: "",
			LastName:  "",
			Body:      "Hi {{.TITLE}} {{.FIRST_NAME}} {{.LAST_NAME}},\nToday, {{.TODAY}}, we would like to tell you that... \nSincrely,\nThe Marketing Team",
			Valid:     false,
		},
		{
			Title:     "Mrrrr",
			FirstName: "ABCD",
			LastName:  "SSH",
			Body:      "Hi Mrrrr ABCD SSH,\nToday, Mon 27 Jun 2022, we would like to tell you that... \nSincrely,\nThe Marketing Team",
			Valid:     false,
		},
		{
			Title:     "sas",
			FirstName: "sss",
			LastName:  "sssssssssssssffasd",
			Body:      "Hi Mr Asda {{.LAST_NAME}},\nToday, abcd, we would like to tell you that... \nSincrely,\nThe Marketing Team",
			Valid:     false,
		},
		{
			Title:     "Mr Mrs Mr",
			FirstName: "Michael",
			LastName:  "Jackson",
			Body:      "Hi Mr Mrs Mr Michael Jackson,\nToday, 27 Jun 2022, we would like to tell you that... \nSincrely,\nThe Marketing Team",
			Valid:     true,
		},
	}

	for _, tc := range tests {

		mail := entity.NewMail("from", "subject", "text/plain", "body")
		customer := entity.NewCustomer(tc.Title, tc.FirstName, tc.LastName, "hihihahahaa@gmail.com")
		mail.GenerateMailFromTemplate(templateMail, *customer)
		if tc.Body == mail.Body {
			assert.Equal(t, tc.Valid, true)
		} else {
			assert.Equal(t, tc.Valid, false)
		}
	}
}

func TestParseEmailInFrom(t *testing.T) {
	type test struct {
		From       string
		ValidEmail string
	}
	tests := []test{
		{
			From:       "The Marketing Team <marketing@example.com>",
			ValidEmail: "marketing@example.com",
		},
		{
			From:       "marketing@example>",
			ValidEmail: "default@example.com",
		},
		{
			From:       "The Marketing Team",
			ValidEmail: "default@example.com",
		},
		{
			From:       "The Marketing Team abcd@abcd mar@a.com",
			ValidEmail: "mar@a.com",
		},
		{
			From:       "The Marketing Team <quanhuynh1310@gmail.com",
			ValidEmail: "quanhuynh1310@gmail.com",
		},
		{
			From:       "The Marketing Team <validemail1302@gmail.com>",
			ValidEmail: "validemail1302@gmail.com",
		},
	}

	for _, tc := range tests {
		mail := entity.NewMail(tc.From, "subject", "text/plain", "body")
		mail.ParseEmailInFrom()
		assert.Equal(t, mail.From, tc.ValidEmail)
	}
}
