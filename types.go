package nogosari

// main type
type app struct {
	Name           string `yaml:"name"`
	Env            string `yaml:"env"`
	Version        string `yaml:"version"`
	HttpConf       *httpConf
	DbConf         *dbConf
	MailerConf     *mailerConf
	ReqLimiterConf *reqLimiterConf
	LoggerConf     *loggerConf
}

type httpConf struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type dbConf struct {
	Dsn          string `yaml:"dsn"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxIdleTime  string `yaml:"maxIdleTime"`
	Dialect      string `yaml:"dialect"`
}

type mailerConf struct {
	Smtp_host     string `yaml:"smtp_host"`
	Smtp_port     int    `yaml:"smtp_port"`
	Smtp_username string `yaml:"smtp_username"`
	Smtp_password string `yaml:"smtp_password"`
	Sender        string `yaml:"sender"`
	TemplateDir   string `yaml:"templateDir"`
}

type reqLimiterConf struct {
	Enabled bool    `yaml:"enabled"`
	Rps     float64 `yaml:"rps"`
	Burst   int     `yaml:"burst"`
}

type loggerConf struct {
	Level  int8   `yaml:"level"`
	Output string `yaml:"output"`
}

// helper type
type mi map[string]interface{}
