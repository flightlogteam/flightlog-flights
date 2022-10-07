package logfile

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElevationSpeedCalculation(t *testing.T) {
	logRecords := `HFDTE200520
	B1559436117621N01023642EA0061800691
	B1559446117621N01023642EA0061900692
   `

	reader, err := NewIGCLogfileReader(logRecords)
	assert.NoError(t, err, "Creating a reader should not result in an error")

	records, _ := reader.ReadRecords()

	assert.Equal(t, 1.0, records[0].elevationSpeed(records[1]), "Lift should be 1 meter a second")
}

func TestGroundSpeedCalculation(t *testing.T) {
	logRecords := `HFDTE200520
    B1559436117621N01023642EA0061800691
	B1610166116618N01023376EA0009900169
   `

	reader, err := NewIGCLogfileReader(logRecords)

	assert.NoError(t, err, "Creating a reader should not result in an error")

	records, _ := reader.ReadRecords()

	assert.Equal(t, 2.96, records[0].groundSpeed(records[1]), "Groundspeed should be higher")
}

func TestFlightLogCreation(t *testing.T) {
	bytes, err := ioutil.ReadFile("./test_igc.igc")

	assert.NoError(t, err, "Test should be able to read the file")

	recordString := string(bytes)

	reader, _ := NewIGCLogfileReader(recordString)
	records, _ := reader.ReadRecords()

	log := NewFlightStats(records)

	assert.Equal(t, 691, log.MaxAltitude)
	assert.Equal(t, 169, log.MinAltitude)
}
