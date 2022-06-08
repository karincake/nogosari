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

	Nogo.initConfig()
	Nogo.initLogger()
	Nogo.initDb()
	// a.initMailer()
	Nogo.initHttp(routerIn)
}

// func GetAppInfo() string {
// 	temp, err := json.Marshal(map[string]string{
// 		"Name":    Nogo.Name,
// 		"Env":     Nogo.Env,
// 		"Version": Nogo.Version,
// 	})
// 	if err == nil {
// 		return string(temp[:])
// 	}
// 	return ""
// }

// func GetHttpConfInfo() string {
// 	temp, err := json.Marshal(httpConfX)
// 	if err == nil {
// 		return string(temp[:])
// 	}
// 	return ""
// }

// func GetDBConfInfo() string {
// 	temp, err := json.Marshal(dbConfX)
// 	if err == nil {
// 		return string(temp[:])
// 	}
// 	return ""
// }

// func GetMailerConfInfo() string {
// 	temp, err := json.Marshal(dbConfX)
// 	if err == nil {
// 		return string(temp[:])
// 	}
// 	return ""
// }
