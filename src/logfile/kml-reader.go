package logfile

import (
	"encoding/xml"
	"errors"
	"strconv"
	"strings"
	"time"
)

type kmlFile struct {
	XMLName  xml.Name `xml:"kml"`
	Document kmlDocument
}

type kmlDocument struct {
	XMLName   xml.Name `xml:"Document"`
	Placemark kmlPlacemark
	Name      string `xml:"name"`
}

type kmlPlacemark struct {
	XMLName    xml.Name `xml:"Placemark"`
	LineString kmlLineString
	TimeSpan   kmlTimeSpan
}

type kmlLineString struct {
	XMLName     xml.Name `xml:"LineString"`
	Coordinates string   `xml:"coordinates"`
}

type kmlTimeSpan struct {
	XMLName xml.Name  `xml:"TimeSpan"`
	Begin   time.Time `xml:"begin"`
}

type KMLLogfileReader struct {
	startTime  time.Time // needed because all time is relative in the KML world
	rawRecords []string  // Parsed from the XML
}

func NewKMLLogfileReader(fileContent []byte) (*KMLLogfileReader, error) {
	// DONE Try to unmarshal the file
	// DONE Read the start time
	// TODO READ each line
	// TODO Add
	var kml kmlFile
	err := xml.Unmarshal(fileContent, &kml)

	if err != nil {
		return nil, errors.New("Unable to read the KML file properly. Got the following error: " + err.Error())
	}

	rawRecords := make([]string, 0)

	for _, line := range strings.Split(kml.Document.Placemark.LineString.Coordinates, "\n") {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) > 0 {
			rawRecords = append(rawRecords, trimmedLine)
		}
	}

	return &KMLLogfileReader{
		startTime:  kml.Document.Placemark.TimeSpan.Begin,
		rawRecords: rawRecords,
	}, nil
}

func (k *KMLLogfileReader) ReadRecords() ([]LogRecord, error) {
	records := make([]LogRecord, len(k.rawRecords))

	for index, record := range k.rawRecords {
		segments := strings.Split(record, ",")
		lon, err := readCoordinate(segments[0])

		if err != nil {
			continue
		}

		lat, err := readCoordinate(segments[1])

		if err != nil {
			continue
		}

		altitude, err := strconv.ParseInt(segments[2], 10, 16)

		records[index] = LogRecord{
			Latitude:          lat,
			Longitude:         lon,
			Time:              k.startTime.Add(time.Second * time.Duration(index)),
			AltitudePreassure: int(altitude),
			AltitudeGNSS:      int(altitude),
		}

	}
	return records, nil
}

func readCoordinate(coordinate string) (float64, error) {
	degrees, err := strconv.ParseFloat(coordinate[1:], 64)

	if err != nil {
		return 0, errors.New("Unable to read coordinate: " + coordinate + " " + err.Error())
	}

	return degrees, nil
}
