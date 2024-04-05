package device

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	// MIGRATE THE SCHEMA
	db.AutoMigrate(&Device{})

	return Repository{
		db: db,
	}
}
