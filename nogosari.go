package nogosari

/*****************************************************************************\
The idea is to place all the needs for a web application to be running in one
place. All the user to do is set the configuration and the handlers since it
is main properties of a website.
\*****************************************************************************/

import (
	"sync"

	"github.com/julienschmidt/httprouter"
	"github.com/karincake/nogosari/logger"
	// "github.com/karincake/nogosari/mailer"
)

type app struct {
	name    string
	env     string
	version string
	// mailer  mailer.Mailer
	logger logger.Logger

	httpConf struct {
		hst  string
		port int
	}
	dbConf struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	mailerConf struct {
		smtp_host     string
		smtp_port     int
		smtp_username string
		smtp_password string
		sender        string
		templateDir   string
	}
	reqLimiterConf struct {
		enabled bool
		rps     float64
		burst   int
	}
	loggerConf struct {
		level  int8
		output string
	}
}

var wgX sync.WaitGroup

// var app *App
func Run(routerIn *httprouter.Router) {
	a := &app{}
	a.initConfig()
	// a.initLogger()
	// a.initDb()
	// a.initMailer()
	// a.initHttp(routerIn)
}
