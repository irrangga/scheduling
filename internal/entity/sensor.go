package entity

import (
	"time"
)

type Sensor struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeviceId  uint      `json:"device_id"`
	Value     int       `json:"value"`
	Unit      string    `json:"unit"`
	Type      string    `json:"type"`
}

type CreateSensor struct {
	DeviceId uint   `json:"device_id"`
	Value    int    `json:"value"`
	Unit     string `json:"unit"`
	Type     string `json:"type"`
}

type UpdateSensor struct {
	ID       uint   `json:"id"`
	DeviceId uint   `json:"device_id"`
	Value    int    `json:"value"`
	Unit     string `json:"unit"`
	Type     string `json:"type"`
}
