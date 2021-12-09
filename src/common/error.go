package common

import "errors"

var (
	// ErrorServiceMissingParameter used when one or more parameters are missing
	ErrorServiceMissingParameter = errors.New("Parameter is missing")
	// ErrorServiceUserNotActive used when a user does not exist
	ErrorServiceUserNotActive = errors.New("User not active")
	// ErrorServiceNoSuchFlight no such flight exists
	ErrorServiceNoSuchFlight = errors.New("No such flight")
	// ErrorServiceNoSuchLocation when the location does not exist
	ErrorServiceNoSuchLocation = errors.New("No such location")
	// ErrorServiceBadCoordinates when coordinates are wrong
	ErrorServiceBadCoordinates = errors.New("Invalid coordinate set")

	ErrorServiceBadDirectionSet = errors.New("The given wind directions are not valid")
)
