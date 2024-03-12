package initializer

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func ConnectToDB(config *Config) error {
	var err error

	dbUser := config.DB.DBUser
	dbPassword := config.DB.DBPassword
	dbName := config.DB.DBName
	dbPort := config.DB.DBPort
	sslConfig := config.DB.DBSSLConfig
	timeZone := config.DB.DBTimeZone

	var connectionString = fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbUser, dbPassword, dbName, dbPort, sslConfig, timeZone,
	)

	gormConfig := gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	Db, err = gorm.Open(postgres.Open(connectionString), &gormConfig)
	if err != nil {
		return err
	}
	return nil
}

func CloseDB() {
	dbInstance, _ := Db.DB()
	dbInstance.Close()
}
