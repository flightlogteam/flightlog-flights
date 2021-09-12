package location

type Reader interface {
	// GetAllLocations get all locations
	GetAllLocations() ([]Location, error)
	// GetLocationById gets a single location by ID
	GetLocationById(id uint) (*Location, error)
}

type Writer interface {
	// CreateLocation Create location
	CreateLocation(location *Location) (uint, error)
	// UpdateLocation Update location
	UpdateLocation(location *Location) error
}

type Repository interface {
	Reader
	Writer
}

type Service interface {
	Reader
	Writer
}
