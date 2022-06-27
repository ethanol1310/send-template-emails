package entity

import "github.com/ethanol1310/send-template-emails/pkg/helper"

type CustomerInfo struct {
	Title     string `csv:"TITLE"`
	FirstName string `csv:"FIRST_NAME"`
	LastName  string `csv:"LAST_NAME"`
	Email     string `csv:"EMAIL"`
}

func NewCustomer(title string, first string, last string, email string) *CustomerInfo {
	return &CustomerInfo{
		Title:     title,
		FirstName: first,
		LastName:  last,
		Email:     email,
	}
}

func (customerInfo *CustomerInfo) ValidCustomer() bool {
	return helper.ValidEmail(customerInfo.Email)
}
