package flight

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/flightlogteam/flightlog-flights/src/common"
	"github.com/flightlogteam/flightlog-flights/src/location"
	"github.com/flightlogteam/userservice/grpc/userservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// FlightService defines the flightservice
type FlightService struct {
	repository      Repository
	userRepository  userservice.UserServiceClient
	locationService location.Service
	config          common.ServiceConfig
}

// NewService - creates a new service
func NewService(serviceUrl string, servicePort string, repository Repository, config common.ServiceConfig, locationService location.Service) (Service, error) {

	connection, err := dialUserService(serviceUrl, servicePort, !config.IsDevMode())

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to connect to userservice at %v:%v", serviceUrl, servicePort))
	}

	return &FlightService{
		repository:      repository,
		userRepository:  userservice.NewUserServiceClient(connection),
		config:          config,
		locationService: locationService,
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

	device, err := f.userRepository.FlyingDeviceExsists(context.Background(), &userservice.FlyingDeviceExsistsRequest{DeviceId: flight.FlyingDevice})

	if !device.Exsists {
		return "", common.ErrorServiceMissingParameter
	}

	flight.Privacy = PrivacyLevel(user.PrivacyLevel)
	flight.UserID = user.UserId

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
	return flight.Duration > 5
}

func dialUserService(serviceUrl string, port string, secure bool) (*grpc.ClientConn, error) {
	log.Println("the security is set to", secure)
	if secure {
		return grpc.Dial(fmt.Sprintf("%s:%s", serviceUrl, "61226"), grpc.WithTransportCredentials(createCredentials()))
	}
	return grpc.Dial(fmt.Sprintf("%s:%s", serviceUrl, port), grpc.WithInsecure())
}

func createCredentials() credentials.TransportCredentials {
	creds, err := credentials.NewClientTLSFromFile("/etc/certificates/server.crt", "")
	if err != nil {
		log.Fatalf("Unable to start the Gateway due to missing certificates. Generate please: %v", err)
	}
	return creds
}
