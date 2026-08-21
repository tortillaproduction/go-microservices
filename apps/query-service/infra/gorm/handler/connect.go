package handler

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	maxRetries    = 10
	retryInterval = 3 * time.Second
)

func ConnectDB() (*gorm.DB, error) {
	dcs := "root:password@tcp(query_db:3306)/sample_db?charset=utf8mb4&parseTime=True&loc=Local"

	var conn *gorm.DB
	var err error

	// Retry connecting ti the DB, since the MySQL container
	// (and its replication setup) may not be ready yet when
	// this service starts.
	for i := 1; i <= maxRetries; i++ {
		conn, err = gorm.Open(mysql.Open(dcs), &gorm.Config{})
		if err == nil {
			if db, dbErr := conn.DB(); dbErr == nil {
				if pingErr := db.Ping(); pingErr == nil {
					// Connection established successfully.
					break
				} else {
					err = pingErr
				}
			} else {
				err = dbErr
			}
		}

		log.Printf("failed to connect db (attempt %d/%d): %v", i, maxRetries, err)

		// Connection failed, wait and retry
		time.Sleep(retryInterval)
	}

	if err != nil {
		// All retries exhusted.
		return nil, DBErrHandler(err)
	}

	db, err := conn.DB()
	if err != nil {
		return nil, DBErrHandler(err)
	}

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	conn.Logger = conn.Logger.LogMode(logger.Info)

	return conn, nil
}
