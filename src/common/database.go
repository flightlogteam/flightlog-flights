package common

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateGormContext(username string, password string, hostname string, port string, database string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(CreateConnectionString(username, password, hostname, port, database)), &gorm.Config{})
}

// EnsureDatabaseExists creates a connection with the given parameters and creates the given database-name if not present
func EnsureDatabaseExists(username string, password string, hostname string, port string, database string) error {
	connectionWithoutDatabase, err := gorm.Open(mysql.Open(CreateConnectionString(username, password, hostname, port, "")), &gorm.Config{})

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
