package nogosari

type mailerConf struct {
	Smtp_host     string `yaml:"smtp_host"`
	Smtp_port     int    `yaml:"smtp_port"`
	Smtp_username string `yaml:"smtp_username"`
	Smtp_password string `yaml:"smtp_password"`
	Sender        string `yaml:"sender"`
	TemplateDir   string `yaml:"templateDir"`
}
