package nogosari

import (
	"os"

	"github.com/karincake/nogosari/logger"
)

func (a *app) initLogger() {
	Logger = *logger.New(os.Stdout, logger.Level(a.LoggerConf.Level))
}
