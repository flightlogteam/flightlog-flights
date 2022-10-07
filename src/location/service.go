package location

import (
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

	if location.ID <= 0 {
		return common.ErrorServiceMissingParameter
	}

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

func (l *LocationService) GetLoactionExistance(ids []uint) (bool, error) {
	return l.repository.GetLoactionExistance(ids)
}

func (l *LocationService) GetClosestLocation(latitude float64, longitude float64, treshold int) ([]Location, error) {

	location := Location{
		Latitude:  latitude,
		Longitude: longitude,
	}

	return l.repository.GetClosestLocation(location.CoordinatesFromRadius(int64(treshold)))
}

func isValidCoordinateSet(lat float64, lon float64) bool {

	if lat > 90 && lat < -90 {
		return false
	}

	if lon > 180 && lon < -180 {
		return false
	}

	return true
}
