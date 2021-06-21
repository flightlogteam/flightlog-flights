package flight

// Writer describes actions related to writing flights
type Writer interface {
	Create(flight *Flight)
	Update(flight *Flight)
}

type Reader interface {
	FlightById(id string)
	FlightByLocation(id int)
}

// Repository describes the repository interface
type Repository interface {
	Reader
	Writer
}

// Service describes the service interface
type Service interface {
	Reader
	Writer
}
