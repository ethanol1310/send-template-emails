/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"crypto/tls"
	"fmt"
	"path/filepath"

	"github.com/ethanol1310/send-template-emails/go-send-email/internal/entity"
	"github.com/ethanol1310/send-template-emails/go-send-email/internal/usecase/esending"
	"github.com/ethanol1310/send-template-emails/go-send-email/internal/usecase/esending/configer"
	"github.com/ethanol1310/send-template-emails/go-send-email/internal/usecase/esending/repo"
	"github.com/ethanol1310/send-template-emails/go-send-email/pkg/common"
	"github.com/ethanol1310/send-template-emails/go-send-email/pkg/helper"
	"github.com/spf13/cobra"
	"gopkg.in/gomail.v2"
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

		// Check path is valid
		if !helper.FileExists(custormers) || filepath.Ext(custormers) != ".csv" ||
			filepath.Ext(errorPath) != ".csv" ||
			!helper.FileExists(template) || filepath.Ext(template) != ".json" ||
			filepath.Ext(output) != ".json" {
			fmt.Printf("File extension wrong or File is not Exist\n")
			return
		}

		// Init SMTP dialer
		eConfig, _ := configer.LoadConfig("./config")
		d := gomail.NewDialer(eConfig.Smtp_host, eConfig.Smtp_port, eConfig.Smtp_username, eConfig.Smtp_password)
		d.TLSConfig = &tls.Config{InsecureSkipVerify: eConfig.Smtp_tls_verify}
		dialer := repo.New(d)

		// Init Service
		service := &esending.ESendingTemplate{Dialer: dialer}
		esendingInfo := entity.ESendingAutomation{
			TemplatePath:      template,
			CustomerPath:      custormers,
			ErrorCustomerPath: errorPath,
			Output:            output}
		erCode := service.PrepareMailToSend(esendingInfo)

		if !common.IS_SUCCESS(erCode) {
			fmt.Printf("PrepareMailToSend - erCode=%d\n", erCode)
		} else {
			fmt.Printf("PrepareMailToSend - erCode=%d\n", erCode)
		}
		service.Send()
		fmt.Printf("FilePath: \nCustomers: %s\nTemplate: %s\nOutput: %s\nErrorPath: %s\n", custormers, template, output, errorPath)
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
