package send

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/ethanol1310/send-template-emails/go-send-email/send/common"
)

type Mail struct {
	From     string
	Subject  string
	MimeType string
	Body     string
}

func (mail *Mail) ReadTemplate(filePath string) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("failed to open json file: %s, error: %v", filePath, err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
		return
	}

	if err := json.Unmarshal(jsonData, &mail); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
		return
	}
}

func (mail *Mail) GenerateContent(templateBody string, customerInfo CustomerInfo) {
	today := time.Now().Format(common.TIME_FORMAT)
	info := TemplateInfo{customerInfo.Title, customerInfo.FirstName, customerInfo.LastName, today}
	mail.Body = info.ParseString(common.TEMPLATE_NAME, templateBody)
}
