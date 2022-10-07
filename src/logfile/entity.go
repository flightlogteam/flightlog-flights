package logfile

import (
	"math"
	"time"

	"github.com/flightlogteam/flightlog-flights/src/location"
)

type FlightLog struct {
	Log     FlightStats
	Start   *location.Location
	Landing *location.Location
}

type LogRecord struct {
	AltitudePreassure int
	AltitudeGNSS      int
	Time              time.Time
	Latitude          float64 // In degrees
	Longitude         float64 // In degrees
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func (s *LogRecord) differenceInMeters(t LogRecord) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = s.Latitude * math.Pi / 180
	lo1 = s.Longitude * math.Pi / 180
	la2 = t.Latitude * math.Pi / 180
	lo2 = t.Longitude * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

// elevationSpeed returns meters per second between two loglines
func (s *LogRecord) elevationSpeed(t LogRecord) float64 {
	meters := t.AltitudePreassure - s.AltitudePreassure
	seconds := t.Time.Sub(s.Time).Seconds()

	return float64(meters) / seconds
}

// groundSpeed as in meter per second
func (s *LogRecord) groundSpeed(t LogRecord) float64 {
	seconds := t.Time.Sub(s.Time).Seconds()
	return math.Round((s.differenceInMeters(t)/seconds)*100) / 100
}

type FlightStats struct {
	Records     []LogRecord
	MaxAltitude int
	MinAltitude int
	MaxClimb    float64
	MaxSink     float64
	MaxSpeed    float64
	Duration    int
}

func NewFlightStats(records []LogRecord) FlightStats {
	log := FlightStats{
		Records:  records,
		Duration: int(records[len(records)-1].Time.Sub(records[0].Time).Seconds()),
	}

	stack := stack{
		size: 3,
	}

	var groundSpeed, climbSpeed float64

	for i, record := range records {
		if i == 0 {
			log.MaxAltitude = record.AltitudeGNSS
			log.MinAltitude = record.AltitudeGNSS
		}

		stack.push(record)
		groundSpeed, climbSpeed = stack.average()

		if groundSpeed > float64(log.MaxSpeed) {
			log.MaxSpeed = groundSpeed
		}

		if climbSpeed > log.MaxClimb {
			log.MaxClimb = climbSpeed
		}

		if climbSpeed < log.MaxSink {
			log.MaxSink = climbSpeed
		}

		if record.AltitudeGNSS > log.MaxAltitude {
			log.MaxAltitude = record.AltitudeGNSS
		}

		if record.AltitudeGNSS < log.MinAltitude {
			log.MinAltitude = record.AltitudeGNSS
		}
	}

	return log
}

type stack struct {
	size    int
	records []LogRecord
}

func (s *stack) push(record LogRecord) {
	if len(s.records) == s.size {
		s.records = s.records[1:]
	}
	s.records = append(s.records, record)
}

// average returns the average (groundSpeed, climb) in the stack.size
func (s *stack) average() (float64, float64) {
	if len(s.records) == s.size {
		groundSpeed := 0
		climbSpeed := 0
		for i := 0; i < s.size; i++ {

			if (i + 1) < s.size {
				groundSpeed += int(s.records[i].groundSpeed(s.records[i+1]))
				climbSpeed += int(s.records[i].elevationSpeed(s.records[i+1]))
			}
		}

		return float64(groundSpeed) / float64(s.size), float64(climbSpeed) / float64(s.size)
	}
	return 0, 0
}
