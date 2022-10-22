package logfile

type LOGFILE int

const (
	LOGFILE_KML = 1
	LOGFILE_IGC = 2
)

type FileReader interface {
	ReadRecords() ([]LogRecord, error)
}

type Service interface {
	ProcessLogfile(records string, logfileType LOGFILE) (*FlightLog, error)
}
