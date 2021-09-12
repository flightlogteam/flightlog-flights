package location

import (
	"github.com/flightlogteam/flightlog-flights/src/common"
)

// Location describes a location... Just like the name intends
type Location struct {
	common.Timestamp
	ID              uint   `gorm:"primaryKey"`
	Latitude        string `gorm:"column:latitude"`
	Longitude       string `gorm:"column:longitude"`
	Name            string `gorm:"column:name"`
	Description     string `gorm:"column:description"`
	FullDescription string `gorm:"column:fullDescription"`
	Altitude        int    `gorm:"column:altitude"`
	Country         string `gorm:"column:contry"`
	CountryPart     string `gorm:"column:contryPart"`
	City            string `gorm:"column:city"`
}

func (Location) TableName() string {
	return "Location"
}
