package entity

import (
	"bytes"
	"text/template"
)

type TemplateInfo struct {
	TITLE      string
	FIRST_NAME string
	LAST_NAME  string
	TODAY      string
}

func (templateInfo *TemplateInfo) ParseString(templateName string, templateBody string) (result string) {
	t, err := template.New(templateName).Parse(templateBody)
	if err != nil {
		panic(err)
	}

	var ret bytes.Buffer
	err = t.Execute(&ret, templateInfo)
	if err != nil {
		panic(err)
	}

	result = ret.String()
	return result
}
