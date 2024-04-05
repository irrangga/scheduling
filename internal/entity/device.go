package entity

import (
	"time"
)

type Device struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Types     []string  `json:"types"`
}

type CreateDevice struct {
	Name  string   `json:"name"`
	Types []string `json:"types"`
}

type UpdateDevice struct {
	ID    uint     `json:"id"`
	Name  string   `json:"name"`
	Types []string `json:"types"`
}
