package flight

import (
	"time"
)

// Timestamp group of timestamps to be used other entities
type Timestamp struct {
	UpdatedAt *time.Time `gorm:"column:updatedat;type:time"`
	CreatedAt *time.Time `gorm:"column:createdat;type:time"`
	DeletedAt *time.Time `gorm:"column:deletedat;type:time"`
}

// Flight describes one flight
type Flight struct {
	Timestamp
	ID           string       `gorm:"primaryKey;column:id;size:64"`
	UserID       string       `gorm:"column:userid;size:64"`
	StartID      int          `gorm:"column:startid;"`  // TODO: add foreign key
	LandingID    int          `gorm:"column:landingid"` // TODO: add foreign key
	FlyingDevice string       `gorm:"column:flyingdevice;size:64"`
	Privacy      PrivacyLevel `gorm:"column:privacylevel"`
	Description  string       `gorm:"column:description"`
	Duration     int          `gorm:"column:duration"`
	MaxHeight    int          `gorm:"column:maxHeight"`
	MinHeight    int          `gorm:"column:minHeight"`
	Distance     int          `gorm:"column:distance"`
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
