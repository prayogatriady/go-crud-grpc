package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_USER     = "root"
	DB_PASSWORD = "root"
	DB_HOST     = "127.0.0.1"
	DB_PORT     = "3306"
	DB_NAME     = "sawer"
)

func InitMySQL() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = CreateTables(db); err != nil {
		return nil, err
	}

	log.Printf("Database connected to %v", DB_NAME)

	return db, nil
}

func CreateTables(db *gorm.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id BIGINT NOT NULL AUTO_INCREMENT,
		username VARCHAR(50) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		balance INT DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP DEFAULT NULL,
		INDEX (id),
		PRIMARY KEY (id)
	)`

	if err := db.Exec(query).Error; err != nil {
		return err
	}
	return nil
}
