package flight

// Writer describes actions related to writing flights
type Writer interface {
	Create(flight *Flight) (string, error)
	Update(flight *Flight) error
}

type Reader interface {
	FlightById(id string) (*Flight, error)
	FlightByLocation(id int) (*Flight, error)
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
