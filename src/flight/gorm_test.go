package flight_test

import (
	"testing"

	"github.com/flightlogteam/flightlog-flights/src/flight"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createSqliteDatabase() (*flight.MariadbRepo, error) {
	db, _ := gorm.Open(sqlite.Open("test-flight.db"), &gorm.Config{})
	return flight.NewRepository(db)
}

func TestMigrationWorks(t *testing.T) {
	repo, err := createSqliteDatabase()

	assert.NotNil(t, repo)
	assert.Nil(t, err)
}
