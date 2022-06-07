package nogosari

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

func (a *app) initConfig() {
	var cfgPath string
	flag.StringVar(&cfgPath, "config-path", ".", "Configuration path (default=./)")

	viper.SetConfigName("config")
	viper.AddConfigPath(cfgPath)
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		panic(err)
	}

	if err := viper.Unmarshal(a); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		panic(err)
	}

	fmt.Println("Loaded config successfully")
}
