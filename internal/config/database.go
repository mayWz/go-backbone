package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB ...
type DB struct {
	DB *gorm.DB
}

func setUpDBConfiguration(env ENV) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.DatabaseUsername,
		env.DatabasePassword,
		env.DatabaseHost,
		env.DatabasePort,
		env.DatabaseName,
	)

	return dsn

}

func ConnectDB(env ENV) (*DB, error) {
	dsn := setUpDBConfiguration(env)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}

	// For logging database
	// db = db.Debug()

	return &DB{DB: db}, nil
}
