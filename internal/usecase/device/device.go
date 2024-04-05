package employee

import (
	"context"
	"iot/internal/entity"
	"strconv"
)

func (uc Usecase) GetListDevices(ctx context.Context) ([]entity.Device, error) {
	devices, err := uc.repo.GetListDevices(ctx)
	if err != nil {
		return []entity.Device{}, err
	}
	return devices, nil
}

func (uc Usecase) GetDevice(ctx context.Context, id string) (entity.Device, error) {
	device, err := uc.repo.GetDevice(ctx, id)
	if err != nil {
		return entity.Device{}, err
	}
	return device, nil
}

func (uc Usecase) CreateDevice(ctx context.Context, input entity.CreateDevice) (entity.Device, error) {
	device, err := uc.repo.CreateDevice(ctx, input)
	if err != nil {
		return entity.Device{}, err
	}
	return device, nil
}

func (uc Usecase) UpdateDevice(ctx context.Context, input entity.UpdateDevice) (entity.Device, error) {
	device, err := uc.repo.GetDevice(ctx, strconv.FormatUint(uint64(input.ID), 10))
	if err != nil {
		return entity.Device{}, err
	}

	updatedDevice, err := uc.repo.UpdateDevice(ctx, input)
	if err != nil {
		return entity.Device{}, err
	}
	updatedDevice.CreatedAt = device.CreatedAt

	return updatedDevice, nil
}

func (uc Usecase) DeleteDevice(ctx context.Context, id string) error {
	_, err := uc.repo.GetDevice(ctx, id)
	if err != nil {
		return err
	}

	err = uc.repo.DeleteDevice(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
