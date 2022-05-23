package db

import (
	"database/sql"
	"fmt"
	"gosql/configs"
	"time"
)

func ConnectDB() *sql.DB {
	dbConfig := configs.GetDBConfig()
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Database)
	db, err := sql.Open("postgres", dbInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(10 * time.Second)
	return db
}
