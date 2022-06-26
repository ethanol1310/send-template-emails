/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/ethanol1310/send-template-emails/go-send-email/internal/entity"
	"github.com/ethanol1310/send-template-emails/go-send-email/internal/usecase/esending"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Automation sending email with template",
	Long:  `Application read list customers' email from CSV and send contents with template to them`,
	Run: func(cmd *cobra.Command, args []string) {
		custormers, _ := cmd.Flags().GetString("customers")
		template, _ := cmd.Flags().GetString("template")
		output, _ := cmd.Flags().GetString("output")
		errorPath, _ := cmd.Flags().GetString("error")
		// ::TODO
		// Check path is valid
		// Check error
		fmt.Printf("%s - %s - %s - %s\n", custormers, template, output, errorPath)

		// ::TODO
		// Call send email function
		esendingInfo := entity.ESendingAutomation{
			TemplatePath:      template,
			CustomerPath:      custormers,
			ErrorCustomerPath: errorPath,
			Output:            output}
		service := &esending.Service{}
		service.Send(esendingInfo)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
