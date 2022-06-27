package entity

import (
	"bytes"
	"text/template"

	"github.com/ethanol1310/send-template-emails/go-send-email/pkg/common"
)

type TemplateInfo struct {
	TITLE      string
	FIRST_NAME string
	LAST_NAME  string
	TODAY      string
}

func (templateInfo *TemplateInfo) ParseString(templateName string, templateBody string) (result string, erCode int) {
	t, err := template.New(templateName).Parse(templateBody)
	if err != nil {
		return result, common.MKFAIL(common.FAILED)
	}

	var ret bytes.Buffer
	err = t.Execute(&ret, templateInfo)
	if err != nil {
		return result, common.MKFAIL(common.FAILED)
	}

	result = ret.String()
	return result, common.MKSUCCESS()
}
