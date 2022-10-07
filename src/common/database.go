package common

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectToDatabase(connectionString string, retryCount int, sleepTime int) (*gorm.DB, error) {
	var err error
	var db *gorm.DB

	db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	attempt := 1

	for err != nil && attempt <= retryCount {
		log.Printf("Could not connect to the database, will wait for %v and this was attempt %v of %v", sleepTime, attempt, retryCount)
		attempt += 1
		time.Sleep(time.Duration(sleepTime) * time.Second)
		db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	}

	return db, err
}
func CreateGormContext(username string, password string, hostname string, port string, database string) (*gorm.DB, error) {
	var err error
	var db *gorm.DB

	sleepTime := 10
	retryCount := 5

	db, err = gorm.Open(mysql.Open(CreateConnectionString(username, password, hostname, port, database)), &gorm.Config{})
	attempt := 1

	for err != nil && attempt <= retryCount {
		log.Printf("Could not connect to the database, will wait for %v and this was attempt %v of %v", sleepTime, attempt, retryCount)
		attempt += 1
		time.Sleep(time.Duration(sleepTime) * time.Second)
		db, err = gorm.Open(mysql.Open(CreateConnectionString(username, password, hostname, port, database)), &gorm.Config{})
	}

	return db, err
}

// EnsureDatabaseExists creates a connection with the given parameters and creates the given database-name if not present
func EnsureDatabaseExists(username string, password string, hostname string, port string, database string) error {
	connectionWithoutDatabase, err := connectToDatabase(CreateConnectionString(username, password, hostname, port, ""), 5, 10)

	if err != nil {
		return err
	}

	tx := connectionWithoutDatabase.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", database))
	tx.Commit()

	return nil
}

// CreateConnectionString creates a MySql / MariaDB connection string from parameters
// hostname and database can be empty strings
func CreateConnectionString(username string, password string, hostname string, port string, database string) string {
	connectionString := fmt.Sprintf("%v:%v@", username, password)

	if len(hostname) > 0 {
		connectionString += fmt.Sprintf("tcp(%v:%v)", hostname, port)
	}

	if len(database) > 0 {
		connectionString += fmt.Sprintf("/%v", database)
	} else {
		connectionString += "/"
	}

	return connectionString + "?parseTime=true"
}
