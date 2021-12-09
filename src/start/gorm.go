package start

import (
	"gorm.io/gorm"
	"time"
)

type StartRepository struct {
	context *gorm.DB
}

func NewRepository(context *gorm.DB) (*StartRepository, error) {
	repo := StartRepository{
		context: context,
	}
	err := context.AutoMigrate(&Start{})
	return &repo, err
}

func (s *StartRepository) Create(start *Start) (uint, error) {
	timestamp := time.Now()
	start.CreatedAt = &timestamp
	start.UpdatedAt = &timestamp

	result := s.context.Create(start)
	return start.ID, result.Error
}

func (s *StartRepository) Update(start *Start) error {
	timestamp := time.Now()
	start.UpdatedAt = &timestamp
	return s.context.Save(start).Error
}

func (s *StartRepository) Delete(id uint) error {
	return s.context.Delete(&Start{}, id).Error
}
