package main

import (
	"flag"

	"github.com/Dalot/goddd/internal/app/adapter"
	"github.com/Dalot/goddd/internal/app/adapter/mysql"
	"github.com/Dalot/goddd/internal/app/adapter/mysql/migrations"
	"github.com/Dalot/goddd/internal/app/domain/factory"
	"github.com/spf13/viper"
)

func main() {
	r := adapter.Router()
	// for local development
	//viper.SetDefault("PGHOST", "0.0.0.0")
	viper.SetDefault("DB_USERNAME", "root")
	viper.SetDefault("DB_PASSWORD", "secret")
	viper.SetDefault("DB_DATABASE", "code_challenge")
	viper.SetDefault("APP_ENV", "debug")

	viper.BindEnv("DB_DATABASE", "DB_DATABASE")
	viper.BindEnv("DB_USERNAME", "DB_USERNAME")
	viper.BindEnv("DB_PASSWORD", "DB_PASSWORD")
	viper.BindEnv("APP_ENV", "APP_ENV")

	conn := mysql.Connection()
	migrations.Migrate()
	sqlDB, err := conn.DB()
	
	defer sqlDB.Close()
	if err != nil {
		panic(err)
	}

	cmd := flag.String("cmd", "", "")
    flag.Parse()
	env := viper.Get("APP_ENV"); 
	if  env == "debug" {
		if string(*cmd) == "migrate" {
			factory.Seed()
		} else if string(*cmd) == "fresh" {
			
			factory.Wipe(conn)
			migrations.Migrate()
			factory.Seed()
		}

	}
	
	r.Run(":8080")
}
