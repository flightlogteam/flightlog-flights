package flight

import (
	"github.com/flightlogteam/flightlog-flights/src/common"
	"github.com/flightlogteam/flightlog-flights/src/location"
	"github.com/flightlogteam/flightlog-flights/src/start"
)

// Flight describes one flight
type Flight struct {
	common.Timestamp
	ID     string `gorm:"primaryKey;column:id;size:64"`
	UserID string `gorm:"column:userid;size:64"`

	StartID      uint              `gorm:"column:startid"`
	LandingID    uint              `gorm:"column:landingid"`
	Start        start.Start       `gorm:"foreignkey:StartID;references:id"`
	Landing      location.Location `gorm:"foreignkey:LandingID;references:id"`
	FlyingDevice string            `gorm:"column:flyingdevice;size:64"`
	Privacy      PrivacyLevel      `gorm:"column:privacylevel"`
	Description  string            `gorm:"column:description"`
	Duration     int               `gorm:"column:duration"`
	MaxHeight    int               `gorm:"column:maxHeight"`
	MinHeight    int               `gorm:"column:minHeight"`
	Distance     int               `gorm:"column:distance"`
}

// TableName - defines which tablename is used in the database
func (Flight) TableName() string {
	return "Flight"
}

// PrivacyLevel describes the shared status of an entity
type PrivacyLevel int

const (
	// ProtectedLevel describes a level where a user can be seen by anyone who has a flightlog login
	ProtectedLevel PrivacyLevel = 0
	// PrivateLevel describes a level where a user needs to grant access to those who wants access
	PrivateLevel PrivacyLevel = 1
	// PublicLevel describes a user that anyone can watch with or without login
	PublicLevel PrivacyLevel = 2
)
