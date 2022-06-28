package helper_test

import (
	"testing"

	"github.com/ethanol1310/send-template-emails/pkg/helper"
	"gopkg.in/go-playground/assert.v1"
)

func TestValidEmailForAll(t *testing.T) {
	type test struct {
		From       string
		ValidEmail bool
	}
	tests := []test{
		{
			From:       "marketing@example.com",
			ValidEmail: true,
		},
		{
			From:       "marketing@example",
			ValidEmail: true,
		},
		{
			From:       "abcd@.ababa",
			ValidEmail: false,
		},
		{
			From:       "mar@a.com",
			ValidEmail: true,
		},
		{
			From:       "quanhuynh1310@gmail.com",
			ValidEmail: true,
		},
		{
			From:       "abcd2@gmail.com",
			ValidEmail: true,
		},
	}

	for _, tc := range tests {
		_, err := helper.ValidEmailRFC5322(tc.From)
		assert.Equal(t, err, tc.ValidEmail)
	}
}
