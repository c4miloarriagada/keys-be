package domain

import "time"

type Key struct {
	ID          int64
	Name        string
	Description string
	Pass        string
	Alias       string
	CreatedAt   *time.Time
	ValidUntil  *time.Time
}

type KeysRepository interface {
	GetByID(id int64) (*Key, error)
	// Save(keys *Keys) error
	// GetAll() ([]Keys, error)
	// DeleteByID(id int64) error
}
