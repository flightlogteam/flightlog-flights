package logfile

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var kmlBytes, _ = ioutil.ReadFile("./test_kml.KML")

func TestCanUnmarshalXMLFile(t *testing.T) {
	reader, err := NewKMLLogfileReader(kmlBytes)

	assert.NoError(t, err, "Should be able to read a KML file")

	assert.Equal(t, 7, len(reader.rawRecords))

	// StartTime is not empty
	assert.True(t, (time.Time{}).Before(reader.startTime))
}

func TestCanProcessRecords(t *testing.T) {
	reader, _ := NewKMLLogfileReader(kmlBytes)
	records, err := reader.ReadRecords()

	if err != nil {
		assert.NoError(t, err, "Unable to produce records")
	}

	assert.Equal(t, 7, len(records))
	firstRecord := records[0]

	assert.Equal(t, 43.7228650, firstRecord.Latitude)
	assert.Equal(t, 6.9439833, firstRecord.Longitude)
	assert.Equal(t, 1126, firstRecord.AltitudeGNSS)
	assert.Equal(t, 1126, firstRecord.AltitudePreassure)
	assert.Equal(t, reader.startTime, firstRecord.Time)

	secondRecord := records[1]

	assert.Equal(t, reader.startTime.Add(time.Second), secondRecord.Time)
}
