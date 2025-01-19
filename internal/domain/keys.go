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

type KeyRepository interface {
	GetByID(id int64) (*Key, error)
	Save(keys *Key) error
	// GetAll() ([]Key, error)
	// DeleteByID(id int64) error
}
