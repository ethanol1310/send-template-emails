package send

import "net/mail"

type CustomerInfo struct {
	Title     string `csv:"TITLE"`
	FirstName string `csv:"FIRST_NAME"`
	LastName  string `csv:"LAST_NAME"`
	Email     string `csv:"EMAIL"`
}

func (customerInfo *CustomerInfo) ValidCustomer() bool {
	return customerInfo.ValidEmail()
}

func (customerInfo *CustomerInfo) ValidEmail() bool {
	_, err := mail.ParseAddress(customerInfo.Email)
	return err == nil
}
