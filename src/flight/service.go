package flight

import (
	"context"
	"errors"
	"fmt"

	"github.com/flightlogteam/flightlog-flights/src/common"
	"github.com/flightlogteam/flightlog-flights/src/location"
	"github.com/flightlogteam/userservice/grpc/userservice"
	"google.golang.org/grpc"
)

// FlightService defines the flightservice
type FlightService struct {
	repository      Repository
	userRepository  userservice.UserServiceClient
	locationService location.Service
}

// NewService - creates a new service
func NewService(serviceUrl string, servicePort string, repository Repository) (Service, error) {

	connection, err := dialUserService(serviceUrl, servicePort)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to connect to userservice at %v:%v", serviceUrl, servicePort))
	}

	return &FlightService{
		repository:     repository,
		userRepository: userservice.NewUserServiceClient(connection),
	}, nil

}

func (f *FlightService) FlightById(id string) (*Flight, error) {
	if len(id) == 0 {
		return nil, common.ErrorServiceMissingParameter
	}

	return f.repository.FlightById(id)
}

func (f *FlightService) FlightByLocation(id int) (*Flight, error) {
	panic("not implemented") // TODO: Implement
}

func (f *FlightService) Create(flight *Flight) (string, error) {

	user, err := f.userRepository.UserByUserId(context.Background(), &userservice.UserByIdRequest{UserId: flight.UserID})

	if err != nil {
		return "", common.ErrorServiceMissingParameter
	}

	if !user.Active {
		return "", common.ErrorServiceUserNotActive
	}

	if !f.validateFlight(flight) {
		return "", common.ErrorServiceMissingParameter
	}

	if !f.validateLocations(flight) {
		return "", common.ErrorServiceNoSuchLocation
	}
	return f.repository.Create(flight)
}

func (f *FlightService) Update(flight *Flight) error {
	existing, err := f.repository.FlightById(flight.ID)

	if err != nil {
		return common.ErrorServiceNoSuchFlight
	}

	if existing == nil {
		return common.ErrorServiceNoSuchFlight
	}

	if !f.validateFlight(flight) {
		return common.ErrorServiceMissingParameter
	}

	if !f.validateLocations(flight) {
		return common.ErrorServiceNoSuchLocation
	}

	return f.repository.Update(flight)
}

func (f *FlightService) validateLocations(flight *Flight) bool {

	locations := []uint{flight.StartID, flight.LandingID}

	for _, id := range locations {
		if loc, err := f.locationService.GetLocationById(id); err != nil || loc == nil {
			return false
		}

	}
	return true
}

func (f *FlightService) validateFlight(flight *Flight) bool {
	return flight.StartID > 0 && flight.LandingID > 0 && len(flight.Description) > 10
}

func dialUserService(serviceUrl string, port string) (*grpc.ClientConn, error) {
	return grpc.Dial(fmt.Sprintf("%s:%s", serviceUrl, port), grpc.WithInsecure())
	//return grpc.Dial(fmt.Sprintf("%s:%s", serviceUrl, "61226"), grpc.WithTransportCredentials(createCredentials()))
}
