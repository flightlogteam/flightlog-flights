package location

type Reader interface {
	// GetAllLocations get all locations
	GetAllLocations() ([]Location, error)
	// GetLocationById gets a single location by ID
	GetLocationById(id uint) (*Location, error)
	GetLoactionExistance(ids []uint) (bool, error)
	// GetClosestLocation gets the closest locations
	// ordered by closest to furthest away
	// `threshold` is the amount of meters
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
	GetClosestLocation(latitude_north float64, latitude_south float64, longitude_east float64, longitude_west float64) ([]Location, error)
}

type Service interface {
	Reader
	Writer
	GetClosestLocation(latitude float64, longitude float64, treshold int) ([]Location, error)
}
