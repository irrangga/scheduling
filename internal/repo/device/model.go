package device

import (
	"iot/internal/repo/sensor"
	"time"

	"gorm.io/datatypes"
)

type Device struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Types     datatypes.JSON
	Sensor    []sensor.Sensor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
