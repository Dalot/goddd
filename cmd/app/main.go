package main

import (
	"github.com/Dalot/goddd/internal/app/adapter"
	"github.com/Dalot/goddd/internal/app/adapter/mysql"
	"github.com/Dalot/goddd/internal/app/adapter/mysql/migrations"
	"github.com/spf13/viper"
)

func main() {
	r := adapter.Router()
	// for local development
	//viper.SetDefault("PGHOST", "0.0.0.0")
	viper.SetDefault("DB_USERNAME", "root")
	viper.SetDefault("DB_PASSWORD", "secret")
	viper.SetDefault("DB_DATABASE", "code_challenge")

	viper.BindEnv("DB_DATABASE", "DB_DATABASE")
	viper.BindEnv("DB_USERNAME", "DB_USERNAME")
	viper.BindEnv("DB_PASSWORD", "DB_PASSWORD")

	conn := mysql.Connection()
	migrations.Migrate()

	r.Run(":8080")
	sqlDB, err := conn.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
}
