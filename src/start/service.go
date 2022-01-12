package start

import (
	"log"

	"github.com/flightlogteam/flightlog-flights/src/common"
	"github.com/flightlogteam/flightlog-flights/src/location"
)

type StartService struct {
	repository      Repository
	locationService location.Service
}

func NewService(repository Repository, locationService location.Service) Service {
	return &StartService{
		repository:      repository,
		locationService: locationService,
	}
}

func (s *StartService) Create(start *Start) (uint, error) {

	for _, l := range append(start.Landings, start.StartLocation) {
		_, err := s.locationService.CreateLocation(&l)
		if err != nil {
			return 0, err
		}
	}

	err := s.validateStart(start)

	log.Printf("Could not validate the start. %v", err)

	if err != nil {
		return 0, err
	}

	return s.repository.Create(start)
}

func (s *StartService) Update(start *Start) error {
	err := s.validateStart(start)

	if err != nil {
		return err
	}

	for _, l := range append(start.Landings, start.StartLocation) {
		s.locationService.CreateLocation(&l)
	}

	return s.repository.Update(start)
}

func (s *StartService) Delete(id uint) error {
	return s.repository.Delete(id)
}

func (s *StartService) validateStart(start *Start) error {
	err := s.validateLoactions(start)

	if err != nil {
		return err
	}

	if !(validateDirection(start.OptimalDirections) && validateDirection(start.SuboptimalDirections)) {
		return common.ErrorServiceBadDirectionSet
	}

	return nil
}

func validateDirection(direction int) bool {
	return direction < 0xFFFFFF && direction >= 0
}

func (s *StartService) validateLoactions(start *Start) error {
	var locationIdList []uint

	locationIdList = append(locationIdList, start.StartLocation.ID)

	for _, landing := range start.Landings {
		locationIdList = append(locationIdList, landing.ID)
	}

	exists, err := s.locationService.GetLoactionExistance(locationIdList)

	if err != nil && !exists {
		return common.ErrorServiceNoSuchLocation
	}

	return nil
}
