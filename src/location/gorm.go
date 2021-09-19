package location

import (
	"time"

	"gorm.io/gorm"
)

type Repo struct {
	context *gorm.DB
}

// NewRepository creates a new repository
func NewRepository(context *gorm.DB) (*Repo, error) {

	repo := Repo{}

	// TODO: Rewrite so all migrations are done from the same location
	err := context.AutoMigrate(&Location{})
	repo.context = context
	return &repo, err
}

func (l *Repo) CreateLocation(location *Location) (uint, error) {
	timestamp := time.Now()
	location.CreatedAt = &timestamp
	location.UpdatedAt = &timestamp
	result := l.context.Create(location)
	return location.ID, result.Error
}

func (l *Repo) UpdateLocation(location *Location) error {
	return l.context.Save(location).Error
}

func (l *Repo) GetAllLocations() ([]Location, error) {
	var locations []Location
	response := l.context.Find(&locations)
	return locations, response.Error
}

func (l *Repo) GetLocationById(id uint) (*Location, error) {
	panic("kfdsndsfnfds")
}
