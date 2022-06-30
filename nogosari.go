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

type app struct {
	Name            string `yaml:"name"`
	Env             string `yaml:"env"`
	Version         string `yaml:"version"`
	HttpConf        *httpConf
	DbConf          *dbConf
	MailerConf      *mailerConf
	RateLimiterConf *rateLimiterConf
	LoggerConf      *loggerConf
}

// export package vars
var Nogo *app
var DB *gorm.DB
var Logger logger.Logger

// internal package vars, for a simpler access
var httpConfX httpConf
var dbConfX dbConf
var mailerConfX mailerConf
var loggerConfX loggerConf

// var app *App
func Run(routerIn *httprouter.Router) {
	Nogo = &app{}
	Nogo.HttpConf = &httpConfX
	Nogo.DbConf = &dbConfX
	Nogo.MailerConf = &mailerConfX
	Nogo.LoggerConf = &loggerConfX

	// like to call manually to make it clear
	Nogo.initConfig()
	Nogo.initLogger()
	Nogo.initDb()
	// a.initMailer()
	Nogo.initHttp(routerIn)
}
