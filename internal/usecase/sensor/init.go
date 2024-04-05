package sensor

import (
	"context"
	"iot/internal/entity"
)

type RepoInterface interface {
	GetListSensors(ctx context.Context) ([]entity.Sensor, error)
	GetSensor(ctx context.Context, id string) (entity.Sensor, error)
	CreateSensor(ctx context.Context, input entity.CreateSensor) (entity.Sensor, error)
	UpdateSensor(ctx context.Context, input entity.UpdateSensor) (entity.Sensor, error)
	DeleteSensor(ctx context.Context, id string) error
}

type Usecase struct {
	repo RepoInterface
}

func New(repo RepoInterface) Usecase {
	return Usecase{
		repo: repo,
	}
}
