package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	goLoggerGormLogger "github.com/pobyzaarif/go-logger/database/framework/gorm/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	Config struct {
		AppMainPort         string
		AppMaxRowPerPage    string
		AppUserPasswordSalt string
		AppJWTSign          string
		DBDriver            string
		// mysql
		DBMySQLHost     string
		DBMySQLPort     string
		DBMySQLUser     string
		DBMySQLPassword string
		DBMySQLName     string
		// sqlite
		DBSQLiteName string
	}
)

func GetAPPConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	return &Config{
		AppMainPort:         os.Getenv("POSLITE_APP_MAIN_PORT"),
		AppMaxRowPerPage:    os.Getenv("POSLITE_APP_MAX_ROW_PER_PAGE"),
		AppUserPasswordSalt: os.Getenv("POSLITE_APP_USER_PASS_SALT"),
		AppJWTSign:          os.Getenv("POSLITE_APP_JWT_SIGN"),
		DBDriver:            os.Getenv("POSLITE_DB_DRIVER"),
		DBMySQLHost:         os.Getenv("POSLITE_DB_MYSQL_HOST"),
		DBMySQLPort:         os.Getenv("POSLITE_DB_MYSQL_PORT"),
		DBMySQLUser:         os.Getenv("POSLITE_DB_MYSQL_USER"),
		DBMySQLPassword:     os.Getenv("POSLITE_DB_MYSQL_PASS"),
		DBMySQLName:         os.Getenv("POSLITE_DB_MYSQL_NAME"),
		DBSQLiteName:        os.Getenv("POSLITE_DB_SQLITE_NAME"),
	}
}

func (conf *Config) GetDatabaseConnection() *gorm.DB {
	if conf.DBDriver == "mysql" {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.DBMySQLUser,
			conf.DBMySQLPassword,
			conf.DBMySQLHost,
			conf.DBMySQLPort,
			conf.DBMySQLName,
		)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newDBLogger()})
		if err != nil {
			log.Fatal(err)
		}

		return db.Debug()
	}

	if conf.DBDriver == "sqlite" {
		db, err := gorm.Open(sqlite.Open(conf.DBSQLiteName), &gorm.Config{Logger: goLoggerGormLogger.Default})
		if err != nil {
			log.Fatal(err)
		}

		return db.Debug()
	}

	log.Fatal("unsupported driver")

	return nil
}

func newDBLogger() logger.Interface {
	return logger.New(
		log.Default(),
		logger.Config{
			SlowThreshold:             30 * time.Second, // Slow SQL threshold
			LogLevel:                  logger.Silent,    // Log level
			IgnoreRecordNotFoundError: false,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,            // Enable Color
		},
	)

}
