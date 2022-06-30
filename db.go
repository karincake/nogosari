package nogosari

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbConf struct {
	Dsn          string `yaml:"dsn"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxIdleTime  string `yaml:"maxIdleTime"`
	Dialect      string `yaml:"dialect"`
}

func (a *app) initDb() {
	if (a.DbConf.Dialect == "") || (a.DbConf.Dsn == "") {
		return
	}
	var gormD gorm.Dialector
	if a.DbConf.Dialect == "mysql" {
		gormD = postgres.Open(a.DbConf.Dsn)
	} else if a.DbConf.Dialect == "postgres" {
		gormD = mysql.Open(a.DbConf.Dsn)
	}

	if db, err := gorm.Open(gormD); err != nil {
		panic("Failed to connect to database!")
	} else {
		fmt.Println("Connected to database: " + db.Name())
		DB = db
	}
}
