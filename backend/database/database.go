package database

import (
	"os"

	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := &mysqldriver.Config{
		DBName:    os.Getenv("DB_NAME"),
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Net:       "tcp",
		Addr:      os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		ParseTime: true,
	}

	dsn := config.FormatDSN()

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database!")
	}

	log.Info().Msg("Database connected!")
}
