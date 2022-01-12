package start

import (
	"github.com/flightlogteam/flightlog-flights/src/common"
	"github.com/flightlogteam/flightlog-flights/src/location"
)

type Start struct {
	common.Timestamp
	ID                   uint                `gorm:"primaryKey"`
	OptimalDirections    int                 `gorm:"column:optimaldirection"`
	SuboptimalDirections int                 `gorm:"column:suboptimaldirection;size:64"`
	StartLocationID      uint                `gorm:"column:startlocationid"`
	StartLocation        location.Location   `gorm:"foreignkey:StartLocationID;references:id"`
	Landings             []location.Location `gorm:"many2many:start_landing_locations"`
}

// TableName - defines which tablename is used in the database
func (Start) TableName() string {
	return "Start"
}
