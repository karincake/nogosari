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
	CodeName        string
	FullName        string
	Env             string
	Version         string
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

// app starter
func Run(appCodeName string, routerIn *httprouter.Router) {
	Nogo = &app{}
	Nogo.CodeName = appCodeName
	Nogo.HttpConf = &httpConfX
	Nogo.DbConf = &dbConfX
	Nogo.MailerConf = &mailerConfX
	Nogo.LoggerConf = &loggerConfX

	// like to call manually to make it clear of what's happening
	Nogo.initConfig()
	Nogo.initLogger()
	Nogo.initDb()
	// a.initMailer()
	Nogo.initHttp(routerIn)
}
