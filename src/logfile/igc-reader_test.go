package logfile

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testFile string = `
AXSBDRP:385831333137030B000C00
HFDTE200520
HFPLTPILOTINCHARGE:
HFCM2CREW2:NIL
HFGTYGLIDERTYPE:
HFGIDGLIDERID:
HFDTMGPSDATUM:WGS84
HFRFWFIRMWAREVERSION:build 5070
HFRHWHARDWAREVERSION:drop_1506
HFFTYFRTYPE:SkyBean,SkyDrop
HFGPSRECEIVER:Quectel,L80,22cm,18000m
HFPRSPRESSALTSENSOR:Measurement specialties,MS5611,25907m
HFALGALTGPS:GEO
HFALPALTPRESSURE:ISA
HFTZNTIMEZONE:+2.0
B1559436117621N01023642EA0061800691
B1559446117621N01023642EA0061800691
B1559456117621N01023642EA0061800691
B1559466117621N01023642EA0061800691
B1559476117621N01023642EA0061800691
B1559486117621N01023642EA0061800691
B1559496117621N01023642EA0061800691
B1559506117621N01023642EA0061800691
B1559516117621N01023642EA0061800691
B1559526117621N01023642EA0061800691
B1559536117621N01023642EA0061800691
B1559546117621N01023642EA0061800691
B1559556117621N01023642EA0061800691
B1559566117621N01023642EA0061800691
`

func TestReadingBRecords(t *testing.T) {
	line := logline("B1559436117621N01023642EA0061800691")

	record, err := line.ReadBRecord(nil)
	expectedTime := time.Time{}.Add(time.Hour * 15).Add(time.Minute * 59).Add(time.Second * 43)

	assert.Equal(t, expectedTime, record.Time, "Time should be parsed correctly")
	assert.NoError(t, err, "There should be no error when parsing the record")

	assert.Equal(t, 61.29368, record.Latitude)
	assert.Equal(t, 10.39403, record.Longitude)
	assert.Equal(t, 618, record.AltitudePreassure, "Altitude preassure should be read correctly")
	assert.Equal(t, 691, record.AltitudeGNSS, "AltitudeGNSS should be read correctly")

}

func TestCreatingIGCLogReader(t *testing.T) {
	fileReader, err := NewIGCLogfileReader(testFile)

	assert.NoError(t, err, "Creating a reader should not result in an error")
	assert.Equal(t, "Europe/Oslo", fileReader.location.String(), "Should be able to parse the correct location")
	assert.Equal(t, 20, fileReader.date.Day(), "Should read days correctly")
	assert.Equal(t, 5, int(fileReader.date.Month()), "Should read months correctly")
	assert.Equal(t, 2020, fileReader.date.Year(), "Should read years correctly")

	records, err := fileReader.ReadRecords()

	assert.NoError(t, err, "Reading of records should not result in an error")

	assert.Equal(t, 14, len(records))
}

func TestCanReadCompleteFile(t *testing.T) {
	bytes, err := ioutil.ReadFile("./test_igc.igc")

	assert.NoError(t, err, "Test should be able to read the file")

	recordString := string(bytes)

	fileReader, err := NewIGCLogfileReader(recordString)
	assert.NoError(t, err, "Creating a reader should not result in an error")

	records, err := fileReader.ReadRecords()

	assert.NoError(t, err, "Records should be produced without an error")

	assert.Equal(t, 629, len(records), "Expected a certain amount of records from that file")
}
