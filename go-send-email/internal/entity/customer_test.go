package entity_test

import (
	"testing"

	"github.com/ethanol1310/send-template-emails/go-send-email/internal/entity"
	"gopkg.in/go-playground/assert.v1"
)

func TestCustomerValidate(t *testing.T) {
	type test struct {
		Title     string
		FirstName string
		LastName  string
		Email     string
		Valid     bool
	}
	tests := []test{
		{
			Title:     "Mr",
			FirstName: "A",
			LastName:  "B",
			Email:     "abacd123",
			Valid:     false,
		},
		{
			Title:     "Mr",
			FirstName: "A",
			LastName:  "B",
			Email:     "abacd123@",
			Valid:     false,
		},
		{
			Title:     "Mr",
			FirstName: "A",
			LastName:  "B",
			Email:     "abacd123@abcd",
			Valid:     true,
		},
		{
			Title:     "Mr",
			FirstName: "A",
			LastName:  "B",
			Email:     "abacd123@gmail.com",
			Valid:     true,
		},
	}
	for _, tc := range tests {

		customer := entity.NewCustomer(tc.Title, tc.FirstName, tc.LastName, tc.Email)
		assert.Equal(t, tc.Valid, customer.ValidCustomer())
	}

}
