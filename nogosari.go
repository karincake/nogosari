package nogosari

/*****************************************************************************\
The idea is to place all the needs for a web application to be running in one
place. All the user to do is set the configuration and the handlers since it
is main properties of a website.
\*****************************************************************************/

import (
	"github.com/julienschmidt/httprouter"
	"github.com/karincake/nogosari/logger"
	"gorm.io/gorm"
	// "github.com/karincake/nogosari/mailer"
)

// main type
type app struct {
	DB *gorm.DB
	// Mailer  mailer.Mailer
	Logger logger.Logger

	Name     string `yaml:"name"`
	Env      string `yaml:"env"`
	Version  string `yaml:"version"`
	HttpConf struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
	DbConf struct {
		Dsn          string `yaml:"dsn"`
		MaxOpenConns int    `yaml:"maxOpenConns"`
		MaxIdleConns int    `yaml:"maxIdleConns"`
		MaxIdleTime  string `yaml:"maxIdleTime"`
		Dialect      string `yaml:"dialect"`
	}
	MailerConf struct {
		Smtp_host     string `yaml:"smtp_host"`
		Smtp_port     int    `yaml:"smtp_port"`
		Smtp_username string `yaml:"smtp_username"`
		Smtp_password string `yaml:"smtp_password"`
		Sender        string `yaml:"sender"`
		TemplateDir   string `yaml:"templateDir"`
	}
	ReqLimiterConf struct {
		Enabled bool    `yaml:"enabled"`
		Rps     float64 `yaml:"rps"`
		Burst   int     `yaml:"burst"`
	}
	LoggerConf struct {
		Level  int8   `yaml:"level"`
		Output string `yaml:"output"`
	}
}

// helper type
type mi map[string]interface{}

// Export the app
var Nogo *app

// var app *App
func Run(routerIn *httprouter.Router) {
	Nogo = &app{}
	Nogo.initConfig()
	Nogo.initLogger()
	Nogo.initDb()
	// a.initMailer()
	Nogo.initHttp(routerIn)
}
