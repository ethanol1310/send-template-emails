package configer

import (
	"github.com/ethanol1310/send-template-emails/go-send-email/pkg/common"
	"github.com/spf13/viper"
)

type Configer struct {
	Template_name        string
	Template_time_format string
	Smtp_host            string
	Smtp_port            int
	Smtp_username        string
	Smtp_password        string
	Smtp_tls_verify      bool
}

func LoadConfig(path string) (configer Configer, erCode int) {
	viper.AddConfigPath(path)
	viper.SetConfigName("esending") // Register config file name (no extension)
	viper.SetConfigType("yaml")     // Look for specific type
	err := viper.ReadInConfig()
	if err != nil {
		return configer, common.MKFAIL(common.NOT_READ)
	}
	configer.Template_name = viper.GetString("template.name")
	configer.Template_time_format = viper.GetString("template.time_format")

	configer.Smtp_host = viper.GetString("smtp.host")
	configer.Smtp_port = viper.GetInt("smtp.port")
	configer.Smtp_username = viper.GetString("smtp.username")
	configer.Smtp_password = viper.GetString("smtp.password")
	configer.Smtp_tls_verify = viper.GetBool("smtp.tls_verify")
	return configer, common.MKSUCCESS()
}
