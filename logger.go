package nogosari

import (
	"os"

	"github.com/karincake/nogosari/logger"
)

type loggerConf struct {
	Level  int8   `yaml:"level"`
	Output string `yaml:"output"`
}

func (a *app) initLogger() {
	Logger = *logger.New(os.Stdout, logger.Level(a.LoggerConf.Level))
}
