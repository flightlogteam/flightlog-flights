package logfile

import (
	"errors"

	"github.com/flightlogteam/flightlog-flights/src/location"
)

func NewLogfileService(locationService location.Service) Service {
	return &LogfileService{
		locationSerice: locationService,
	}
}

type LogfileService struct {
	locationSerice location.Service
}

func (l *LogfileService) ProcessLogfile(records string, logfileType LOGFILE) (*FlightLog, error) {

	var reader FileReader
	var err error

	switch logfileType {
	case LOGFILE_KML:
		reader, err = NewKMLLogfileReader([]byte(records))
		break
	case LOGFILE_IGC:
		reader, err = NewIGCLogfileReader(records)
		break
	default:
		return nil, errors.New("Not a supported logfile")
	}

	if err != nil {
		return nil, err
	}

	flightRecords, err := reader.ReadRecords()

	flightStats := NewFlightStats(flightRecords)

	startRecord := flightStats.Records[0]
	landingRecord := flightStats.Records[len(flightStats.Records)-1]

	starts, err := l.locationSerice.GetClosestLocation(startRecord.Latitude, startRecord.Longitude, 100)

	var start, landing *location.Location = nil, nil

	if err == nil && len(starts) > 0 {
		start = &starts[0]
	}

	landings, err := l.locationSerice.GetClosestLocation(landingRecord.Latitude, landingRecord.Longitude, 100)

	if err == nil && len(landings) > 0 {
		landing = &landings[0]
	}

	return &FlightLog{
		Log:     flightStats,
		Start:   start,
		Landing: landing,
	}, nil

}
