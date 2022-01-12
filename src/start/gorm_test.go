package start_test

import (
	"os"
	"testing"

	"github.com/flightlogteam/flightlog-flights/src/location"
	"github.com/flightlogteam/flightlog-flights/src/start"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbname = "test-start.db"

func createSqliteDatabase() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	return db
}

func removeSqliteDatabaseFile() error {
	return os.Remove(dbname)
}

func createStartRepo(db *gorm.DB) (start.Repository, error) {
	return start.NewRepository(db)
}

func createLocationRepo(db *gorm.DB) (location.Repository, error) {
	return location.NewRepository(db)
}

func TestEnsureRepoCreation(t *testing.T) {
	db := createSqliteDatabase()
	defer removeSqliteDatabaseFile()
	repo, err := createStartRepo(db)

	assert.NotNil(t, repo)
	assert.Nil(t, err)
}

func TestStartCreation(t *testing.T) {
	database := createSqliteDatabase()
	defer removeSqliteDatabaseFile()
	locationRepo, _ := createLocationRepo(database)
	startRepo, _ := createStartRepo(database)

	testLocation := location.Location{
		Latitude:        "61",
		Longitude:       "10",
		Name:            "TestLoc1",
		Description:     "Some short description",
		FullDescription: "Full description",
		Country:         "Norway",
		CountryPart:     "Innlandet",
		City:            "Oslo",
	}

	_, err := locationRepo.CreateLocation(&testLocation)

	assert.Nil(t, err)

	testStart := start.Start{
		OptimalDirections:    1,
		SuboptimalDirections: 3,
		StartLocation:        testLocation,
		Landings:             make([]location.Location, 0),
	}

	testStart.Landings = append(testStart.Landings, testLocation)

	_, err = startRepo.Create(&testStart)

	assert.Nil(t, err)

	assert.NotEqual(t, 0, testStart.ID, "The start gets an ID from the database that is not 0")

	// Ensure that the many to many relation has been resolved in the database

	rows, err := database.Raw("SELECT * FROM start_landing_locations").Rows()

	assert.Nil(t, err)

	assert.Equal(t, true, rows.Next())
	assert.Equal(t, false, rows.Next())
}
