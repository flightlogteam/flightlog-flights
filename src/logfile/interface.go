package logfile

type FileReader interface {
	ReadRecords() ([]LogRecord, error)
}

type Service interface {
	ProcessIGCLogfile(records string) (*FlightLog, error)
}
