package location_test

import (
	"os"
	"testing"

	"github.com/flightlogteam/flightlog-flights/src/location"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbname = "test-location.db"

var testLocation = location.Location{
	Latitude:        "61",
	Longitude:       "10",
	Name:            "TestLoc1",
	Description:     "Some short description",
	FullDescription: "Full description",
	Country:         "Norway",
	CountryPart:     "Innlandet",
	City:            "Oslo",
}

func createSqliteDatabase() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	return db
}

func removeSqliteDatabaseFile() error {
	return os.Remove(dbname)
}

func createLocationRepo(db *gorm.DB) (location.Repository, error) {
	return location.NewRepository(db)
}

func TestLocationCreation(t *testing.T) {
	database := createSqliteDatabase()
	locationRepo, err := createLocationRepo(database)
	defer removeSqliteDatabaseFile()

	assert.Nil(t, err)

	id, err := locationRepo.CreateLocation(&testLocation)

	assert.Positive(t, id, "The new location should have a positive ID")
	assert.Nil(t, err)
}

func TestLocationExsists(t *testing.T) {
	database := createSqliteDatabase()
	locationRepo, _ := createLocationRepo(database)
	defer removeSqliteDatabaseFile()

	id, _ := locationRepo.CreateLocation(&testLocation)

	exists, err := locationRepo.GetLoactionExistance([]uint{id})

	assert.Nil(t, err)
	assert.True(t, exists, "The location should exist")
}

func TestGetAllLocations(t *testing.T) {
	database := createSqliteDatabase()
	locationRepo, _ := createLocationRepo(database)
	defer removeSqliteDatabaseFile()

	locationRepo.CreateLocation(&testLocation)
	testLocation.ID = 0
	locationRepo.CreateLocation(&testLocation)
	testLocation.ID = 0
	locationRepo.CreateLocation(&testLocation)

	locations, err := locationRepo.GetAllLocations()

	assert.Nil(t, err)
	assert.Exactly(t, 3, len(locations), "Should be exactly 3 locations in the database")
}

func TestGetLocationById(t *testing.T) {
	database := createSqliteDatabase()
	locationRepo, _ := createLocationRepo(database)
	defer removeSqliteDatabaseFile()

	id, _ := locationRepo.CreateLocation(&testLocation)

	location, err := locationRepo.GetLocationById(id)

	assert.Nil(t, err)

	assert.Equal(t, testLocation.Name, location.Name)
	assert.Equal(t, testLocation.Altitude, location.Altitude)
	assert.Equal(t, testLocation.Description, location.Description)
	assert.Equal(t, testLocation.FullDescription, location.FullDescription)
}
