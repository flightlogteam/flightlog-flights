package flight

import "gorm.io/gorm"

type MariadbRepo struct {
	context *gorm.DB
}
