package logfile

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/bradfitz/latlong"
)

type logline string
type recordType string

const (
	BRecord       recordType = "B"
	UnknownRecord recordType = "Unknown"
)

func (line logline) length() int {
	return len(strings.TrimSpace(string(line)))
}

func (line logline) RecordType() recordType {
	lineString := strings.TrimSpace(string(line))

	if len(lineString) > 0 && lineString[0] == 'B' {
		return BRecord
	}
	return UnknownRecord
}

// 0					 10				   20				   30
// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4
// B 1 5 5 9 4 3 6 1 1 7 6 2 1 N 0 1 0 2 3 6 4 2 E A 0 0 6 1 8 0 0 6 9 1
func (line logline) ReadBRecord(date *time.Time) (LogRecord, error) {
	line = logline(strings.TrimSpace(string(line)))
	if len(line) != 35 {
		return LogRecord{}, errors.New("Invalid record")
	}

	if date == nil {
		date = &time.Time{}
	}
	time, err := parseTime(string(line[1:7]), *date)

	latitude, err := parseCoordinate(string(line[7:14]))
	longitude, err := parseCoordinate(string(line[16:23]))

	altitudePreassure, _ := strconv.ParseInt(string(line[25:30]), 10, 32)
	altitudeGNSS, _ := strconv.ParseInt(string(line[30:35]), 10, 32)

	return LogRecord{
		AltitudePreassure: int(altitudePreassure),
		AltitudeGNSS:      int(altitudeGNSS),
		Latitude:          latitude,
		Longitude:         longitude,
		Time:              time,
	}, err
}

func NewIGCLogfileReader(records string) (*IGCLogfileReader, error) {
	rawRecords := strings.Split(records, "\n")

	bRecord, err := findFirstBrecord(rawRecords)

	if err != nil {
		return nil, errors.New("Could not find a single BRecord")
	}

	record, err := logline(bRecord).ReadBRecord(nil)

	if err != nil {
		return nil, errors.New("Unable to read this file. Not able to resolve location." + err.Error())
	}

	reader := &IGCLogfileReader{
		rawRecords: rawRecords,
	}

	reader.resolveLocation(record.Latitude, record.Longitude)
	dateIndex := strings.Index(records, "HFDTE") + 5
	rawDate := records[dateIndex : dateIndex+6]
	reader.readDate(rawDate)

	return reader, nil
}

type IGCLogfileReader struct {
	location   time.Location
	rawRecords []string
	date       time.Time
}

func findFirstBrecord(records []string) (string, error) {
	index := 0
	var record string
	for index < len(records) {
		record = records[index]
		if logline(record).RecordType() == BRecord {
			return record, nil
		}
		index += 1
	}
	return "", errors.New("No BRecords in this file")
}

func (r *IGCLogfileReader) readDate(date string) {
	day, _ := strconv.ParseInt(date[:2], 10, 16)
	month, _ := strconv.ParseInt(date[2:4], 10, 16)
	year, _ := strconv.ParseInt(date[4:6], 10, 16)
	year += 2000

	r.date = time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, &r.location)
}

func (r *IGCLogfileReader) resolveLocation(latitude float64, longitude float64) error {
	timeZone := latlong.LookupZoneName(latitude, longitude)
	location, err := time.LoadLocation(timeZone)
	r.location = *location
	return err
}

func (r *IGCLogfileReader) ReadRecords() ([]LogRecord, error) {
	var records []LogRecord
	for _, segment := range r.rawRecords {
		line := logline(segment)
		if line.RecordType() == BRecord && line.length() == 35 {
			logRecord, _ := line.ReadBRecord(&r.date)
			records = append(records, logRecord)
		}
	}

	return records, nil
}

// parseCoordinate reads one set of coordinate and returns a float
func parseCoordinate(coordinate string) (float64, error) {
	deg, err := strconv.ParseInt(coordinate[:2], 10, 16)

	if err != nil {
		return 0, errors.New("Unable to parse the coordinate")
	}

	minutes, err := strconv.ParseInt(coordinate[2:], 10, 16)

	if err != nil {
		return 0, errors.New("Unable to parse the coordinate")
	}

	return math.Round((float64(deg)+float64(minutes)/60000)*100000) / 100000, nil

}

// parseTime formatted as HHMMSS
func parseTime(timeString string, date time.Time) (time.Time, error) {
	hours, err := strconv.ParseInt(timeString[:2], 10, 8)

	if err != nil {
		return time.Time{}, err
	}

	minutes, err := strconv.ParseInt(timeString[2:4], 10, 8)

	if err != nil {
		return time.Time{}, err
	}
	seconds, err := strconv.ParseInt(timeString[4:6], 10, 8)

	if err != nil {
		return time.Time{}, err
	}

	return time.Date(date.Day(), date.Month(), date.Year(), int(hours), int(minutes), int(seconds), 0, date.Location()), nil
}
