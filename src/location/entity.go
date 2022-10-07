package location

import (
	"math"
	"sort"

	"github.com/flightlogteam/flightlog-flights/src/common"
)

// Location describes a location... Just like the name intends
type Location struct {
	common.Timestamp
	ID              uint    `gorm:"primaryKey"`
	Latitude        float64 `gorm:"column:latitude"`
	Longitude       float64 `gorm:"column:longitude"`
	Name            string  `gorm:"column:name"`
	Description     string  `gorm:"column:description"`
	FullDescription string  `gorm:"column:fullDescription"`
	Altitude        int     `gorm:"column:altitude"`
	Country         string  `gorm:"column:contry"`
	CountryPart     string  `gorm:"column:contryPart"`
	City            string  `gorm:"column:city"`
}

func (Location) TableName() string {
	return "Location"
}

type Locations []Location

// CoordinatesFromRadius returns coordinates in the
// radius arround the location
// `radius` is in meters
// returns (lat-north, lat-south, long-west, long-east)
func (l *Location) CoordinatesFromRadius(radius int64) (float64, float64, float64, float64) {
	earth_radius := 6378000.0 // in meters
	latitudeDegrees := (float64(radius) / earth_radius) * (180.0 / math.Pi)
	longitudeDegrees := (float64(radius) / float64(earth_radius)) * (180.0 / math.Pi) / math.Cos(l.Latitude*math.Pi/180.0)

	return l.Latitude + latitudeDegrees, l.Latitude - latitudeDegrees, l.Longitude - longitudeDegrees, l.Latitude + longitudeDegrees
}

func (l Locations) OrderByLocation(latitude float64, longitude float64) {
	sort.Slice(l, func(i, j int) bool {
		return l[i].coordinateDiff(latitude, longitude) < l[j].coordinateDiff(latitude, longitude)
	})
}

func (l *Location) coordinateDiff(lat float64, lon float64) float64 {
	latDiff := l.Latitude - lat

	if latDiff < 0 {
		latDiff = latDiff * -1
	}

	lonDiff := l.Longitude - lon

	if lonDiff < 0 {
		lonDiff = lonDiff * -1
	}

	return latDiff + lonDiff
}
