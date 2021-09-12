package location

import (
	"strconv"

	"github.com/flightlogteam/flightlog-flights/src/common"
)

type LocationService struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &LocationService{
		repository: repo,
	}
}

// GetAllLocations get all locations
func (l *LocationService) GetAllLocations() ([]Location, error) {
	return l.repository.GetAllLocations()
}

// GetLocationById gets a single location by ID
func (l *LocationService) GetLocationById(id uint) (*Location, error) {
	return l.GetLocationById(id)
}

// CreateLocation Create location
func (l *LocationService) CreateLocation(location *Location) (uint, error) {
	err := l.validateLocation(location)

	if err != nil {
		return 0, err
	}

	return l.repository.CreateLocation(location)
}

// UpdateLocation Update location
func (l *LocationService) UpdateLocation(location *Location) error {
	err := l.validateLocation(location)

	if err != nil {
		return err
	}

	return l.repository.UpdateLocation(location)
}

func (l *LocationService) validateLocation(location *Location) error {
	if !isValidCoordinateSet(location.Latitude, location.Longitude) {
		return common.ErrorServiceBadCoordinates
	}

	if len(location.City) < 3 && len(location.Country) < 3 && len(location.CountryPart) < 3 {
		return common.ErrorServiceMissingParameter
	}

	return nil
}

func isValidCoordinateSet(lat string, lon string) bool {
	latNumber, err := strconv.ParseFloat(lat, 32)
	if err != nil {
		return false
	}

	if latNumber > 90 && latNumber < -90 {
		return false
	}

	lonNumber, err := strconv.ParseFloat(lon, 32)
	if err != nil {
		return false
	}

	if lonNumber > 180 && latNumber < -180 {
		return false
	}

	return true
}
