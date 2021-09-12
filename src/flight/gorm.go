package flight

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MariadbRepo struct holding
type MariadbRepo struct {
	context *gorm.DB
}

// NewRepository creates a new repository
func NewRepository(context *gorm.DB) (*MariadbRepo, error) {

	repo := MariadbRepo{}

	err := context.AutoMigrate(&Flight{})
	repo.context = context
	return &repo, err
}

func (m *MariadbRepo) Create(flight *Flight) (string, error) {
	flight.ID = uuid.New().String()

	timestamp := time.Now()

	flight.CreatedAt = &timestamp
	flight.UpdatedAt = &timestamp

	result := m.context.Create(flight)
	return flight.ID, result.Error
}

func (m *MariadbRepo) Update(flight *Flight) error {
	return m.context.Save(flight).Error
}

func (m *MariadbRepo) FlightById(id string) (*Flight, error) {
	var flight Flight
	result := m.context.Where("id=?", id).First(&flight)
	return &flight, result.Error
}

func (m *MariadbRepo) FlightByLocation(id int) (*Flight, error) {
	panic("not implemented") // TODO: Implement
}
