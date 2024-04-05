package sensor

import (
	"context"
	"iot/internal/entity"
)

type UcInterface interface {
	GetListSensors(ctx context.Context) ([]entity.Sensor, error)
	GetSensor(ctx context.Context, id string) (entity.Sensor, error)
	CreateSensor(ctx context.Context, input entity.CreateSensor) (entity.Sensor, error)
	UpdateSensor(ctx context.Context, input entity.UpdateSensor) (entity.Sensor, error)
	DeleteSensor(ctx context.Context, id string) error
}

type Handler struct {
	uc UcInterface
}

func New(uc UcInterface) Handler {
	return Handler{
		uc: uc,
	}
}
