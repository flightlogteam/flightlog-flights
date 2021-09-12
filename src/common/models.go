package common

import "time"

type Timestamp struct {
	UpdatedAt *time.Time `gorm:"column:updatedat;type:time"`
	CreatedAt *time.Time `gorm:"column:createdat;type:time"`
	DeletedAt *time.Time `gorm:"column:deletedat;type:time"`
}
