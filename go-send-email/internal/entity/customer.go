package entity

import "github.com/ethanol1310/send-template-emails/go-send-email/pkg/helper"

type CustomerInfo struct {
	Title     string `csv:"TITLE"`
	FirstName string `csv:"FIRST_NAME"`
	LastName  string `csv:"LAST_NAME"`
	Email     string `csv:"EMAIL"`
}

func (customerInfo *CustomerInfo) ValidCustomer() bool {
	return helper.ValidEmail(customerInfo.Email)
}
