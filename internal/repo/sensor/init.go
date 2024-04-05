package sensor

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	// MIGRATE THE SCHEMA
	db.AutoMigrate(&Sensor{})

	return Repository{
		db: db,
	}
}
