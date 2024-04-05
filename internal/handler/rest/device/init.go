package device

import (
	"context"
	"iot/internal/entity"
)

type UcInterface interface {
	GetListDevices(ctx context.Context) ([]entity.Device, error)
	GetDevice(ctx context.Context, id string) (entity.Device, error)
	CreateDevice(ctx context.Context, input entity.CreateDevice) (entity.Device, error)
	UpdateDevice(ctx context.Context, input entity.UpdateDevice) (entity.Device, error)
	DeleteDevice(ctx context.Context, id string) error
}

type Handler struct {
	uc UcInterface
}

func New(uc UcInterface) Handler {
	return Handler{
		uc: uc,
	}
}
