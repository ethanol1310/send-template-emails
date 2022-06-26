package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/mail"
	"os"
	"path/filepath"
)

func WriteStringToFile(filePath string, content string, append bool) {
	var flag int
	if append {
		flag = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	} else {
		flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	}

	file, err := os.OpenFile(filePath, flag, 0644)
	if err != nil {
		log.Println(err)
	}

	defer file.Close()
	if _, err := file.WriteString(content); err != nil {
		log.Println(err)
	}
}

func WriteDataToJson(filePath string, content interface{}) {
	if filepath.Ext(filePath) != ".json" {
		log.Println("filePath is not json")
		return
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	dataBytes, err := json.MarshalIndent(content, "", "\t")
	if err != nil {
		log.Println(err)
	}

	if _, err := fmt.Fprintf(file, "%s\n", dataBytes); err != nil {
		fmt.Println(err)
	}
}

func FileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
