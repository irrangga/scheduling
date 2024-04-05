package sensor

import (
	"context"
	"iot/internal/entity"
	"time"
)

func (r Repository) GetListSensors(ctx context.Context) ([]entity.Sensor, error) {
	var sensors []Sensor

	err := r.db.WithContext(ctx).Find(&sensors).Error
	if err != nil {
		return []entity.Sensor{}, err
	}

	var sensorsEntity []entity.Sensor

	for _, sensor := range sensors {
		sensorsEntity = append(sensorsEntity, entity.Sensor{
			ID:        sensor.ID,
			CreatedAt: sensor.CreatedAt,
			UpdatedAt: sensor.UpdatedAt,
			DeviceId:  sensor.DeviceId,
			Value:     sensor.Value,
			Unit:      sensor.Unit,
			Type:      sensor.Type,
		})
	}
	return sensorsEntity, nil
}

func (r Repository) GetSensor(ctx context.Context, id string) (entity.Sensor, error) {
	var sensor Sensor

	err := r.db.WithContext(ctx).First(&sensor, id).Error
	if err != nil {
		return entity.Sensor{}, err
	}

	return entity.Sensor{
		ID:        sensor.ID,
		CreatedAt: sensor.CreatedAt,
		UpdatedAt: sensor.UpdatedAt,
		DeviceId:  sensor.DeviceId,
		Value:     sensor.Value,
		Unit:      sensor.Unit,
		Type:      sensor.Type,
	}, nil
}

func (r Repository) CreateSensor(ctx context.Context, input entity.CreateSensor) (entity.Sensor, error) {
	sensor := Sensor{
		DeviceId: input.DeviceId,
		Value:    input.Value,
		Unit:     input.Unit,
		Type:     input.Type,
	}

	err := r.db.WithContext(ctx).Create(&sensor).Error
	if err != nil {
		return entity.Sensor{}, err
	}

	return entity.Sensor{
		ID:        sensor.ID,
		CreatedAt: sensor.CreatedAt,
		UpdatedAt: sensor.UpdatedAt,
		DeviceId:  sensor.DeviceId,
		Value:     sensor.Value,
		Unit:      sensor.Unit,
		Type:      sensor.Type,
	}, nil
}

func (r Repository) UpdateSensor(ctx context.Context, input entity.UpdateSensor) (entity.Sensor, error) {
	sensor := Sensor{
		ID:       input.ID,
		DeviceId: input.DeviceId,
		Value:    input.Value,
		Unit:     input.Unit,
		Type:     input.Type,
	}

	err := r.db.WithContext(ctx).Model(&sensor).Updates(&Sensor{
		UpdatedAt: time.Now(),
		DeviceId:  input.DeviceId,
		Value:     input.Value,
		Unit:      input.Unit,
		Type:      input.Type,
	}).Error

	if err != nil {
		return entity.Sensor{}, err
	}

	return entity.Sensor{
		ID:        sensor.ID,
		CreatedAt: sensor.CreatedAt,
		UpdatedAt: sensor.UpdatedAt,
		DeviceId:  sensor.DeviceId,
		Value:     sensor.Value,
		Unit:      sensor.Unit,
		Type:      sensor.Type,
	}, nil
}

func (r Repository) DeleteSensor(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Delete(&Sensor{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
