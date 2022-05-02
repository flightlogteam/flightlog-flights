package logfile

type FileReader interface {
	ReadRecords() ([]logRecord, error)
}
