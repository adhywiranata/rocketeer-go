package db

import (
	"database/sql"
	"fmt"
)

// DBStruct to ensure database configuration
type DBStruct struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// ConnectDB to connect to database
func ConnectDB(port, host, user, password, name string) (*sql.DB, error) {
	dbConfig := DBStruct{
		Port:     port,
		Host:     host,
		User:     user,
		Password: password,
		Name:     name,
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
