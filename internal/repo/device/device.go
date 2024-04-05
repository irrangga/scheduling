package device

import (
	"context"
	"encoding/json"
	"iot/internal/entity"
	"time"
)

func (r Repository) GetListDevices(ctx context.Context) ([]entity.Device, error) {
	var devices []Device

	err := r.db.WithContext(ctx).Find(&devices).Error
	if err != nil {
		return []entity.Device{}, err
	}

	var devicesEntity []entity.Device

	for _, device := range devices {
		var typesEntity []string
		err = json.Unmarshal(device.Types, &typesEntity)
		if err != nil {
			return []entity.Device{}, err
		}

		devicesEntity = append(devicesEntity, entity.Device{
			ID:        device.ID,
			CreatedAt: device.CreatedAt,
			UpdatedAt: device.UpdatedAt,
			Name:      device.Name,
			Types:     typesEntity,
		})
	}
	return devicesEntity, nil
}

func (r Repository) GetDevice(ctx context.Context, id string) (entity.Device, error) {
	var device Device

	err := r.db.WithContext(ctx).First(&device, id).Error
	if err != nil {
		return entity.Device{}, err
	}

	var typesEntity []string
	err = json.Unmarshal(device.Types, &typesEntity)
	if err != nil {
		return entity.Device{}, err
	}

	return entity.Device{
		ID:        device.ID,
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.UpdatedAt,
		Name:      device.Name,
		Types:     typesEntity,
	}, nil
}

func (r Repository) CreateDevice(ctx context.Context, input entity.CreateDevice) (entity.Device, error) {
	typesJson, err := json.Marshal(input.Types)
	if err != nil {
		return entity.Device{}, err
	}

	device := Device{
		Name:  input.Name,
		Types: typesJson,
	}

	err = r.db.WithContext(ctx).Create(&device).Error
	if err != nil {
		return entity.Device{}, err
	}

	var typesEntity []string
	err = json.Unmarshal(typesJson, &typesEntity)
	if err != nil {
		return entity.Device{}, err
	}

	return entity.Device{
		ID:        device.ID,
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.UpdatedAt,
		Name:      device.Name,
		Types:     typesEntity,
	}, nil
}

func (r Repository) UpdateDevice(ctx context.Context, input entity.UpdateDevice) (entity.Device, error) {
	typesJson, err := json.Marshal(input.Types)
	if err != nil {
		return entity.Device{}, err
	}

	device := Device{
		ID:    input.ID,
		Name:  input.Name,
		Types: typesJson,
	}

	var typesEntity []string
	err = json.Unmarshal(typesJson, &typesEntity)
	if err != nil {
		return entity.Device{}, err
	}

	err = r.db.WithContext(ctx).Model(&device).Updates(&Device{
		UpdatedAt: time.Now(),
		Name:      device.Name,
		Types:     device.Types,
	}).Error

	if err != nil {
		return entity.Device{}, err
	}

	return entity.Device{
		ID:        device.ID,
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.UpdatedAt,
		Name:      device.Name,
		Types:     typesEntity,
	}, nil
}

func (r Repository) DeleteDevice(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Delete(&Device{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
