package sensor

import (
	"context"
	"iot/internal/entity"
	"strconv"
)

func (uc Usecase) GetListSensors(ctx context.Context) ([]entity.Sensor, error) {
	sensors, err := uc.repo.GetListSensors(ctx)
	if err != nil {
		return []entity.Sensor{}, err
	}
	return sensors, nil
}

func (uc Usecase) GetSensor(ctx context.Context, id string) (entity.Sensor, error) {
	sensor, err := uc.repo.GetSensor(ctx, id)
	if err != nil {
		return entity.Sensor{}, err
	}
	return sensor, nil
}

func (uc Usecase) CreateSensor(ctx context.Context, input entity.CreateSensor) (entity.Sensor, error) {
	sensor, err := uc.repo.CreateSensor(ctx, input)
	if err != nil {
		return entity.Sensor{}, err
	}
	return sensor, nil
}

func (uc Usecase) UpdateSensor(ctx context.Context, input entity.UpdateSensor) (entity.Sensor, error) {
	sensor, err := uc.repo.GetSensor(ctx, strconv.FormatUint(uint64(input.ID), 10))
	if err != nil {
		return entity.Sensor{}, err
	}

	updatedSensor, err := uc.repo.UpdateSensor(ctx, input)
	if err != nil {
		return entity.Sensor{}, err
	}
	updatedSensor.CreatedAt = sensor.CreatedAt

	return updatedSensor, nil
}

func (uc Usecase) DeleteSensor(ctx context.Context, id string) error {
	_, err := uc.repo.GetSensor(ctx, id)
	if err != nil {
		return err
	}

	err = uc.repo.DeleteSensor(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
