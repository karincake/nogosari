package nogosari

import (
	"fmt"

	"github.com/spf13/viper"
)

func (a *app) initConfig() {
	viper.SetConfigName("api")
	viper.AddConfigPath("../../conf")
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

	fmt.Println("Config is loaded successfully")
}
