package sensor

import (
	"time"
)

type Sensor struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeviceId  uint
	Value     int
	Unit      string
	Type      string
}
