package flight

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MariadbRepo struct holding
type MariadbRepo struct {
	context *gorm.DB
}

// NewRepository creates a new repository
func NewRepository(username string, password string, hostname string, port string, database string) (*MariadbRepo, error) {
	err := ensureDatabaseExists(username, password, hostname, port, database)

	if err != nil {
		return nil, err
	}

	repo := MariadbRepo{}
	context, err := gorm.Open(mysql.Open(createConnectionString(username, password, hostname, port, database)), &gorm.Config{})

	context.AutoMigrate(User{}, &Credentials{})
	repo.context = context
	return &repo, err
}

func (m *MariadbRepo) Create(flight *flight.Flight) {
	panic("not implemented") // TODO: Implement
}

func (m *MariadbRepo) Update(flight *flight.Flight) {
	panic("not implemented") // TODO: Implement
}

func (m *MariadbRepo) FlightById(id string) {
	panic("not implemented") // TODO: Implement
}

func (m *MariadbRepo) FlightByLocation(id int) {
	panic("not implemented") // TODO: Implement
}

func createConnectionString(username string, password string, hostname string, port string, database string) string {
	connectionString := fmt.Sprintf("%v:%v@", username, password)

	if len(hostname) > 0 {
		connectionString += fmt.Sprintf("tcp(%v:%v)", hostname, port)
	}

	if len(database) > 0 {
		connectionString += fmt.Sprintf("/%v", database)
	} else {
		connectionString += "/"
	}

	return connectionString + "?parseTime=true"
}
