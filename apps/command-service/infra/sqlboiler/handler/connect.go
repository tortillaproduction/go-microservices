package handler

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/aarondl/sqlboiler/v4/boil"
)

type DBConfig struct {
	Dbname string `toml:"dbname"`
	Host   string `toml:"host"`
	Port   int64  `toml:"port"`
	User   string `toml:"user"`
	Pass   string `toml:"pass"`
}

const (
	maxRetries    = 10
	retryInterval = 3 * time.Second
)

// tomlRead reads DB settings from `database.toml` and returns DBConfig type.
func tomlRead() (*DBConfig, error) {
	path := os.Getenv("DATABASE_TOML_PATH")
	if path == "" {
		path = "infra/sqlboiler/config/database.toml"
	}

	m := map[string]DBConfig{}
	_, err := toml.DecodeFile(path, &m)
	if err != nil {
		return nil, err
	}
	config := m["mysql"]

	// Override if MYSQL_HOST environment variable is present.
	if envHost := os.Getenv("MYSQL_HOST"); envHost != "" {
		config.Host = envHost
	}

	return &config, nil
}

// DBConnect connect to the database.
func DBConnect() error {
	config, err := tomlRead()
	if err != nil {
		return DBErrHandler(err)
	}

	rdbms := "mysql"
	connect_str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Pass, config.Host, config.Port, config.Dbname)

	var conn *sql.DB

	// Retry connecting ti the DB, since the MySQL container
	// may not be ready yet when this service starts.
	for i := 0; i <= maxRetries; i++ {
		conn, err = sql.Open(rdbms, connect_str)
		if err == nil {
			if pingErr := conn.Ping(); pingErr == nil {
				// Connection established successfully.
				break
			} else {
				err = pingErr
			}
		}

		log.Printf("failed to connect db (attempt %d/%d): %v", i, maxRetries, err)

		time.Sleep(retryInterval)
	}

	if err != nil {
		// All retries exhausted.
		return DBErrHandler(err)
	}

	MAX_IDLE_CONNS := 10
	MAX_OPEN_CONNS := 100
	CONN_MAX_LIFETIME := 300 * time.Second

	// Set connection pool.
	conn.SetMaxIdleConns(MAX_IDLE_CONNS)
	conn.SetMaxOpenConns(MAX_OPEN_CONNS)
	conn.SetConnMaxLifetime(CONN_MAX_LIFETIME)

	boil.SetDB(conn)
	boil.DebugMode = true
	return nil
}
