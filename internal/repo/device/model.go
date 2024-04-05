package device

import (
	"time"

	"gorm.io/datatypes"
)

type Device struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Types     datatypes.JSON
}
